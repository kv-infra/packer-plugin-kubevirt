// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

package iso

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ptr "k8s.io/utils/ptr"

	v1 "kubevirt.io/api/core/v1"
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

func virtualMachine(name, isoVolumeName, diskSize, instanceType, preferenceName string) *v1.VirtualMachine {
	cdrom := uint(1)
	oemdrv := uint(2)
	rootdisk := uint(3)

	return &v1.VirtualMachine{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "VirtualMachine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name + "-vm",
		},
		Spec: v1.VirtualMachineSpec{
			RunStrategy: ptr.To(v1.RunStrategyAlways),
			Instancetype: &v1.InstancetypeMatcher{
				Name: instanceType,
			},
			Preference: &v1.PreferenceMatcher{
				Name: preferenceName,
			},
			DataVolumeTemplates: []v1.DataVolumeTemplateSpec{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: name + "-rootdisk",
					},
					Spec: cdiv1.DataVolumeSpec{
						PVC: &corev1.PersistentVolumeClaimSpec{
							Resources: corev1.VolumeResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(diskSize),
								},
							},
							AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
						},
						Source: &cdiv1.DataVolumeSource{
							Blank: &cdiv1.DataVolumeBlankImage{},
						},
					},
				},
			},
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				Spec: v1.VirtualMachineInstanceSpec{
					Domain: v1.DomainSpec{
						Devices: v1.Devices{
							Disks: []v1.Disk{
								{
									Name: "cdrom",
									DiskDevice: v1.DiskDevice{
										CDRom: &v1.CDRomTarget{
											Tray: "closed",
										},
									},
									BootOrder: &cdrom,
								},
								{
									Name: "oemdrv",
									DiskDevice: v1.DiskDevice{
										CDRom: &v1.CDRomTarget{
											Tray: "closed",
										},
									},
									BootOrder: &oemdrv,
								},
								{
									Name: "rootdisk",
									DiskDevice: v1.DiskDevice{
										Disk: &v1.DiskTarget{
											Bus: "virtio",
										},
									},
									BootOrder: &rootdisk,
								},
							},
						},
					},
					Volumes: []v1.Volume{
						{
							Name: "cdrom",
							VolumeSource: v1.VolumeSource{
								DataVolume: &v1.DataVolumeSource{
									Name: isoVolumeName,
								},
							},
						},
						{
							Name: "rootdisk",
							VolumeSource: v1.VolumeSource{
								DataVolume: &v1.DataVolumeSource{
									Name: name + "-rootdisk",
								},
							},
						},
						{
							Name: "oemdrv",
							VolumeSource: v1.VolumeSource{
								ConfigMap: &v1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "oemdrv-cm",
									},
									VolumeLabel: "OEMDRV",
								},
							},
						},
					},
				},
			},
		},
	}
}

func cloneVolume(name, namespace, diskSize string) *cdiv1.DataVolume {
	return &cdiv1.DataVolume{
		TypeMeta: metav1.TypeMeta{
			APIVersion: cdiv1.CDIGroupVersionKind.GroupVersion().String(),
			Kind:       "DataVolume",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: cdiv1.DataVolumeSpec{
			Source: &cdiv1.DataVolumeSource{
				PVC: &cdiv1.DataVolumeSourcePVC{
					Name:      name + "-rootdisk",
					Namespace: namespace,
				},
			},
			PVC: &corev1.PersistentVolumeClaimSpec{
				Resources: corev1.VolumeResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(diskSize),
					},
				},
				AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			},
		},
	}
}

func sourceVolume(name, namespace, instanceType, preferenceName string) *cdiv1.DataSource {
	return &cdiv1.DataSource{
		TypeMeta: metav1.TypeMeta{
			APIVersion: cdiv1.CDIGroupVersionKind.GroupVersion().String(),
			Kind:       "DataSource",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"instancetype.kubevirt.io/default-instancetype": instanceType,
				"instancetype.kubevirt.io/default-preference":   preferenceName,
			},
		},
		Spec: cdiv1.DataSourceSpec{
			Source: cdiv1.DataSourceSource{
				PVC: &cdiv1.DataVolumeSourcePVC{
					Name:      name,
					Namespace: namespace,
				},
			},
		},
	}
}
