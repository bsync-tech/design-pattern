package bridge

type Brand interface {
	GetBrand() string
}

type HuaweiBrand struct{}

func (h HuaweiBrand) GetBrand() string {
	return "huawei"
}

type AppleBrand struct{}

func (a AppleBrand) GetBrand() string {
	return "apple"
}

type DeviceType interface {
	GetDeviceType() string
}

type PhoneDeviceType struct{}

func (p *PhoneDeviceType) GetDeviceType() string {
	return "phone"
}

type PadDeviceType struct{}

func (p *PadDeviceType) GetDeviceType() string {
	return "pad"
}

type Device struct {
	brand      Brand
	deviceType DeviceType
}

func (d *Device) SetBrand(brand Brand) {
	d.brand = brand
}

func (d *Device) SetDeviceType(deviceType DeviceType) {
	d.deviceType = deviceType
}

func (d *Device) String() string {
	return d.brand.GetBrand() + "-" + d.deviceType.GetDeviceType()
}
