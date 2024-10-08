//go:build !ignore_autogenerated

/*
Copyright The CloudNativePG Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package api

import (
	pkgapi "github.com/cloudnative-pg/machinery/pkg/api"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCredentials) DeepCopyInto(out *AzureCredentials) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
	if in.StorageAccount != nil {
		in, out := &in.StorageAccount, &out.StorageAccount
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
	if in.StorageKey != nil {
		in, out := &in.StorageKey, &out.StorageKey
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
	if in.StorageSasToken != nil {
		in, out := &in.StorageSasToken, &out.StorageSasToken
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCredentials.
func (in *AzureCredentials) DeepCopy() *AzureCredentials {
	if in == nil {
		return nil
	}
	out := new(AzureCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarmanCredentials) DeepCopyInto(out *BarmanCredentials) {
	*out = *in
	if in.Google != nil {
		in, out := &in.Google, &out.Google
		*out = new(GoogleCredentials)
		(*in).DeepCopyInto(*out)
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(S3Credentials)
		(*in).DeepCopyInto(*out)
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarmanCredentials.
func (in *BarmanCredentials) DeepCopy() *BarmanCredentials {
	if in == nil {
		return nil
	}
	out := new(BarmanCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarmanObjectStoreConfiguration) DeepCopyInto(out *BarmanObjectStoreConfiguration) {
	*out = *in
	in.BarmanCredentials.DeepCopyInto(&out.BarmanCredentials)
	if in.EndpointCA != nil {
		in, out := &in.EndpointCA, &out.EndpointCA
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
	if in.Wal != nil {
		in, out := &in.Wal, &out.Wal
		*out = new(WalBackupConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(DataBackupConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HistoryTags != nil {
		in, out := &in.HistoryTags, &out.HistoryTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarmanObjectStoreConfiguration.
func (in *BarmanObjectStoreConfiguration) DeepCopy() *BarmanObjectStoreConfiguration {
	if in == nil {
		return nil
	}
	out := new(BarmanObjectStoreConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataBackupConfiguration) DeepCopyInto(out *DataBackupConfiguration) {
	*out = *in
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = new(int32)
		**out = **in
	}
	if in.AdditionalCommandArgs != nil {
		in, out := &in.AdditionalCommandArgs, &out.AdditionalCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataBackupConfiguration.
func (in *DataBackupConfiguration) DeepCopy() *DataBackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(DataBackupConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleCredentials) DeepCopyInto(out *GoogleCredentials) {
	*out = *in
	if in.ApplicationCredentials != nil {
		in, out := &in.ApplicationCredentials, &out.ApplicationCredentials
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleCredentials.
func (in *GoogleCredentials) DeepCopy() *GoogleCredentials {
	if in == nil {
		return nil
	}
	out := new(GoogleCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Credentials) DeepCopyInto(out *S3Credentials) {
	*out = *in
	if in.AccessKeyIDReference != nil {
		in, out := &in.AccessKeyIDReference, &out.AccessKeyIDReference
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
	if in.SecretAccessKeyReference != nil {
		in, out := &in.SecretAccessKeyReference, &out.SecretAccessKeyReference
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
	if in.RegionReference != nil {
		in, out := &in.RegionReference, &out.RegionReference
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
	if in.SessionToken != nil {
		in, out := &in.SessionToken, &out.SessionToken
		*out = new(pkgapi.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Credentials.
func (in *S3Credentials) DeepCopy() *S3Credentials {
	if in == nil {
		return nil
	}
	out := new(S3Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WalBackupConfiguration) DeepCopyInto(out *WalBackupConfiguration) {
	*out = *in
	if in.ArchiveAdditionalCommandArgs != nil {
		in, out := &in.ArchiveAdditionalCommandArgs, &out.ArchiveAdditionalCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RestoreAdditionalCommandArgs != nil {
		in, out := &in.RestoreAdditionalCommandArgs, &out.RestoreAdditionalCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WalBackupConfiguration.
func (in *WalBackupConfiguration) DeepCopy() *WalBackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(WalBackupConfiguration)
	in.DeepCopyInto(out)
	return out
}
