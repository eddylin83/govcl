
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TFont struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewFont() *TFont {
    f := new(TFont)
    f.instance = Font_Create()
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsFont(obj interface{}) *TFont {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TFont{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsFont.
func FontFromInst(inst uintptr) *TFont {
    return AsFont(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsFont.
func FontFromObj(obj IObject) *TFont {
    return AsFont(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsFont.
func FontFromUnsafePointer(ptr unsafe.Pointer) *TFont {
    return AsFont(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (f *TFont) Free() {
    if f.instance != 0 {
        Font_Free(f.instance)
        f.instance, f.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (f *TFont) Instance() uintptr {
    return f.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (f *TFont) UnsafeAddr() unsafe.Pointer {
    return f.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (f *TFont) IsValid() bool {
    return f.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (f *TFont) Is() TIs {
    return TIs(f.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (f *TFont) As() TAs {
//    return TAs(f.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TFontClass() TClass {
    return Font_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (f *TFont) Assign(Source IObject) {
    Font_Assign(f.instance, CheckPtr(Source))
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (f *TFont) HandleAllocated() bool {
    return Font_HandleAllocated(f.instance)
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (f *TFont) GetNamePath() string {
    return Font_GetNamePath(f.instance)
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (f *TFont) DisposeOf() {
    Font_DisposeOf(f.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (f *TFont) ClassType() TClass {
    return Font_ClassType(f.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (f *TFont) ClassName() string {
    return Font_ClassName(f.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (f *TFont) InstanceSize() int32 {
    return Font_InstanceSize(f.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (f *TFont) InheritsFrom(AClass TClass) bool {
    return Font_InheritsFrom(f.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (f *TFont) Equals(Obj IObject) bool {
    return Font_Equals(f.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (f *TFont) GetHashCode() int32 {
    return Font_GetHashCode(f.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (f *TFont) ToString() string {
    return Font_ToString(f.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (f *TFont) Handle() HFONT {
    return Font_GetHandle(f.instance)
}

// CN: 设置控件句柄。
// EN: Set Control handle.
func (f *TFont) SetHandle(value HFONT) {
    Font_SetHandle(f.instance, value)
}

func (f *TFont) PixelsPerInch() int32 {
    return Font_GetPixelsPerInch(f.instance)
}

func (f *TFont) SetPixelsPerInch(value int32) {
    Font_SetPixelsPerInch(f.instance, value)
}

func (f *TFont) Charset() TFontCharset {
    return Font_GetCharset(f.instance)
}

func (f *TFont) SetCharset(value TFontCharset) {
    Font_SetCharset(f.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (f *TFont) Color() TColor {
    return Font_GetColor(f.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (f *TFont) SetColor(value TColor) {
    Font_SetColor(f.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (f *TFont) Height() int32 {
    return Font_GetHeight(f.instance)
}

// CN: 设置高度。
// EN: Set height.
func (f *TFont) SetHeight(value int32) {
    Font_SetHeight(f.instance, value)
}

// CN: 获取组件名称。
// EN: Get the component name.
func (f *TFont) Name() string {
    return Font_GetName(f.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (f *TFont) SetName(value string) {
    Font_SetName(f.instance, value)
}

func (f *TFont) Orientation() int32 {
    return Font_GetOrientation(f.instance)
}

func (f *TFont) SetOrientation(value int32) {
    Font_SetOrientation(f.instance, value)
}

func (f *TFont) Pitch() TFontPitch {
    return Font_GetPitch(f.instance)
}

func (f *TFont) SetPitch(value TFontPitch) {
    Font_SetPitch(f.instance, value)
}

func (f *TFont) Size() int32 {
    return Font_GetSize(f.instance)
}

func (f *TFont) SetSize(value int32) {
    Font_SetSize(f.instance, value)
}

func (f *TFont) Style() TFontStyles {
    return Font_GetStyle(f.instance)
}

func (f *TFont) SetStyle(value TFontStyles) {
    Font_SetStyle(f.instance, value)
}

func (f *TFont) Quality() TFontQuality {
    return Font_GetQuality(f.instance)
}

func (f *TFont) SetQuality(value TFontQuality) {
    Font_SetQuality(f.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (f *TFont) SetOnChange(fn TNotifyEvent) {
    Font_SetOnChange(f.instance, fn)
}

