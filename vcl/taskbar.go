
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

type TTaskbar struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskbar(owner IComponent) *TTaskbar {
    t := new(TTaskbar)
    t.instance = Taskbar_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsTaskbar(obj interface{}) *TTaskbar {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTaskbar{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsTaskbar.
func TaskbarFromInst(inst uintptr) *TTaskbar {
    return AsTaskbar(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsTaskbar.
func TaskbarFromObj(obj IObject) *TTaskbar {
    return AsTaskbar(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTaskbar.
func TaskbarFromUnsafePointer(ptr unsafe.Pointer) *TTaskbar {
    return AsTaskbar(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TTaskbar) Free() {
    if t.instance != 0 {
        Taskbar_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskbar) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskbar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskbar) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TTaskbar) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TTaskbar) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskbarClass() TClass {
    return Taskbar_StaticClassType()
}

func (t *TTaskbar) DoThumbButtonNotify(ItemID uint16) {
    Taskbar_DoThumbButtonNotify(t.instance, ItemID)
}

func (t *TTaskbar) DoThumbPreviewRequest(APreviewHeight uint16, APreviewWidth uint16) {
    Taskbar_DoThumbPreviewRequest(t.instance, APreviewHeight , APreviewWidth)
}

func (t *TTaskbar) DoWindowPreviewRequest() {
    Taskbar_DoWindowPreviewRequest(t.instance)
}

func (t *TTaskbar) UnregisterTab() {
    Taskbar_UnregisterTab(t.instance)
}

func (t *TTaskbar) ApplyButtonsChanges() {
    Taskbar_ApplyButtonsChanges(t.instance)
}

func (t *TTaskbar) GetMainWindowHwnd() HWND {
    return Taskbar_GetMainWindowHwnd(t.instance)
}

func (t *TTaskbar) GetOverlayHIcon() HICON {
    return Taskbar_GetOverlayHIcon(t.instance)
}

func (t *TTaskbar) ActivateTab() bool {
    return Taskbar_ActivateTab(t.instance)
}

func (t *TTaskbar) InvalidateThumbPreview() {
    Taskbar_InvalidateThumbPreview(t.instance)
}

func (t *TTaskbar) UpdateTab() {
    Taskbar_UpdateTab(t.instance)
}

func (t *TTaskbar) CheckApplyChanges() {
    Taskbar_CheckApplyChanges(t.instance)
}

func (t *TTaskbar) ApplyChanges() {
    Taskbar_ApplyChanges(t.instance)
}

func (t *TTaskbar) ApplyOverlayChanges() {
    Taskbar_ApplyOverlayChanges(t.instance)
}

func (t *TTaskbar) ApplyProgressChanges() {
    Taskbar_ApplyProgressChanges(t.instance)
}

func (t *TTaskbar) ApplyTabsChanges() {
    Taskbar_ApplyTabsChanges(t.instance)
}

func (t *TTaskbar) ApplyClipAreaChanges() {
    Taskbar_ApplyClipAreaChanges(t.instance)
}

func (t *TTaskbar) ClearClipArea() {
    Taskbar_ClearClipArea(t.instance)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTaskbar) FindComponent(AName string) *TComponent {
    return AsComponent(Taskbar_FindComponent(t.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskbar) GetNamePath() string {
    return Taskbar_GetNamePath(t.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTaskbar) HasParent() bool {
    return Taskbar_HasParent(t.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskbar) Assign(Source IObject) {
    Taskbar_Assign(t.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTaskbar) DisposeOf() {
    Taskbar_DisposeOf(t.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskbar) ClassType() TClass {
    return Taskbar_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskbar) ClassName() string {
    return Taskbar_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskbar) InstanceSize() int32 {
    return Taskbar_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskbar) InheritsFrom(AClass TClass) bool {
    return Taskbar_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskbar) Equals(Obj IObject) bool {
    return Taskbar_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskbar) GetHashCode() int32 {
    return Taskbar_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TTaskbar) ToString() string {
    return Taskbar_ToString(t.instance)
}

func (t *TTaskbar) TaskBarButtons() *TThumbBarButtonList {
    return AsThumbBarButtonList(Taskbar_GetTaskBarButtons(t.instance))
}

func (t *TTaskbar) SetTaskBarButtons(value *TThumbBarButtonList) {
    Taskbar_SetTaskBarButtons(t.instance, CheckPtr(value))
}

func (t *TTaskbar) ProgressState() TTaskBarProgressState {
    return Taskbar_GetProgressState(t.instance)
}

func (t *TTaskbar) SetProgressState(value TTaskBarProgressState) {
    Taskbar_SetProgressState(t.instance, value)
}

func (t *TTaskbar) ProgressMaxValue() int64 {
    return Taskbar_GetProgressMaxValue(t.instance)
}

func (t *TTaskbar) SetProgressMaxValue(value int64) {
    Taskbar_SetProgressMaxValue(t.instance, value)
}

func (t *TTaskbar) ProgressValue() int64 {
    return Taskbar_GetProgressValue(t.instance)
}

func (t *TTaskbar) SetProgressValue(value int64) {
    Taskbar_SetProgressValue(t.instance, value)
}

func (t *TTaskbar) OverlayIcon() *TIcon {
    return AsIcon(Taskbar_GetOverlayIcon(t.instance))
}

func (t *TTaskbar) SetOverlayIcon(value *TIcon) {
    Taskbar_SetOverlayIcon(t.instance, CheckPtr(value))
}

func (t *TTaskbar) OverlayHint() string {
    return Taskbar_GetOverlayHint(t.instance)
}

func (t *TTaskbar) SetOverlayHint(value string) {
    Taskbar_SetOverlayHint(t.instance, value)
}

func (t *TTaskbar) PreviewClipRegion() *TPreviewClipRegion {
    return AsPreviewClipRegion(Taskbar_GetPreviewClipRegion(t.instance))
}

func (t *TTaskbar) SetPreviewClipRegion(value *TPreviewClipRegion) {
    Taskbar_SetPreviewClipRegion(t.instance, CheckPtr(value))
}

func (t *TTaskbar) TabProperties() TThumbTabProperties {
    return Taskbar_GetTabProperties(t.instance)
}

func (t *TTaskbar) SetTabProperties(value TThumbTabProperties) {
    Taskbar_SetTabProperties(t.instance, value)
}

func (t *TTaskbar) ToolTip() string {
    return Taskbar_GetToolTip(t.instance)
}

func (t *TTaskbar) SetToolTip(value string) {
    Taskbar_SetToolTip(t.instance, value)
}

func (t *TTaskbar) SetOnThumbPreviewRequest(fn TThumbPreviewItemRequestEvent) {
    Taskbar_SetOnThumbPreviewRequest(t.instance, fn)
}

func (t *TTaskbar) SetOnWindowPreviewItemRequest(fn TWindowPreviewItemRequestEvent) {
    Taskbar_SetOnWindowPreviewItemRequest(t.instance, fn)
}

func (t *TTaskbar) SetOnThumbButtonClick(fn TThumbButtonNotifyEvent) {
    Taskbar_SetOnThumbButtonClick(t.instance, fn)
}

func (t *TTaskbar) TaskbarIsAvailable() bool {
    return Taskbar_GetTaskbarIsAvailable(t.instance)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTaskbar) ComponentCount() int32 {
    return Taskbar_GetComponentCount(t.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (t *TTaskbar) ComponentIndex() int32 {
    return Taskbar_GetComponentIndex(t.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (t *TTaskbar) SetComponentIndex(value int32) {
    Taskbar_SetComponentIndex(t.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTaskbar) Owner() *TComponent {
    return AsComponent(Taskbar_GetOwner(t.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTaskbar) Name() string {
    return Taskbar_GetName(t.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTaskbar) SetName(value string) {
    Taskbar_SetName(t.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTaskbar) Tag() int {
    return Taskbar_GetTag(t.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTaskbar) SetTag(value int) {
    Taskbar_SetTag(t.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTaskbar) Components(AIndex int32) *TComponent {
    return AsComponent(Taskbar_GetComponents(t.instance, AIndex))
}

