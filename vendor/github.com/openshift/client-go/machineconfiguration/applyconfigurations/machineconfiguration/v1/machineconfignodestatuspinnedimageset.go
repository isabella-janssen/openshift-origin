// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// MachineConfigNodeStatusPinnedImageSetApplyConfiguration represents a declarative configuration of the MachineConfigNodeStatusPinnedImageSet type for use
// with apply.
type MachineConfigNodeStatusPinnedImageSetApplyConfiguration struct {
	Name                      *string `json:"name,omitempty"`
	CurrentGeneration         *int32  `json:"currentGeneration,omitempty"`
	DesiredGeneration         *int32  `json:"desiredGeneration,omitempty"`
	LastFailedGeneration      *int32  `json:"lastFailedGeneration,omitempty"`
	LastFailedGenerationError *string `json:"lastFailedGenerationError,omitempty"`
}

// MachineConfigNodeStatusPinnedImageSetApplyConfiguration constructs a declarative configuration of the MachineConfigNodeStatusPinnedImageSet type for use with
// apply.
func MachineConfigNodeStatusPinnedImageSet() *MachineConfigNodeStatusPinnedImageSetApplyConfiguration {
	return &MachineConfigNodeStatusPinnedImageSetApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MachineConfigNodeStatusPinnedImageSetApplyConfiguration) WithName(value string) *MachineConfigNodeStatusPinnedImageSetApplyConfiguration {
	b.Name = &value
	return b
}

// WithCurrentGeneration sets the CurrentGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentGeneration field is set to the value of the last call.
func (b *MachineConfigNodeStatusPinnedImageSetApplyConfiguration) WithCurrentGeneration(value int32) *MachineConfigNodeStatusPinnedImageSetApplyConfiguration {
	b.CurrentGeneration = &value
	return b
}

// WithDesiredGeneration sets the DesiredGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredGeneration field is set to the value of the last call.
func (b *MachineConfigNodeStatusPinnedImageSetApplyConfiguration) WithDesiredGeneration(value int32) *MachineConfigNodeStatusPinnedImageSetApplyConfiguration {
	b.DesiredGeneration = &value
	return b
}

// WithLastFailedGeneration sets the LastFailedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastFailedGeneration field is set to the value of the last call.
func (b *MachineConfigNodeStatusPinnedImageSetApplyConfiguration) WithLastFailedGeneration(value int32) *MachineConfigNodeStatusPinnedImageSetApplyConfiguration {
	b.LastFailedGeneration = &value
	return b
}

// WithLastFailedGenerationError sets the LastFailedGenerationError field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastFailedGenerationError field is set to the value of the last call.
func (b *MachineConfigNodeStatusPinnedImageSetApplyConfiguration) WithLastFailedGenerationError(value string) *MachineConfigNodeStatusPinnedImageSetApplyConfiguration {
	b.LastFailedGenerationError = &value
	return b
}
