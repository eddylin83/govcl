
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Gauge_Create(AOwner: TComponent): TGauge; extdecl;
begin
  Result :=  TGauge.Create(AOwner);
end;

procedure Gauge_Free(AObj: TGauge); extdecl;
begin
  AObj.Free;
end;

procedure Gauge_AddProgress(AObj: TGauge; Value: Integer); extdecl;
begin
  AObj.AddProgress(Value);
end;

procedure Gauge_BringToFront(AObj: TGauge); extdecl;
begin
  AObj.BringToFront;
end;

procedure Gauge_ClientToScreen(AObj: TGauge; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure Gauge_ClientToParent(AObj: TGauge; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function Gauge_Dragging(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.Dragging;
end;

function Gauge_HasParent(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Gauge_Hide(AObj: TGauge); extdecl;
begin
  AObj.Hide;
end;

procedure Gauge_Invalidate(AObj: TGauge); extdecl;
begin
  AObj.Invalidate;
end;

function Gauge_Perform(AObj: TGauge; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Gauge_Refresh(AObj: TGauge); extdecl;
begin
  AObj.Refresh;
end;

procedure Gauge_Repaint(AObj: TGauge); extdecl;
begin
  AObj.Repaint;
end;

procedure Gauge_ScreenToClient(AObj: TGauge; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure Gauge_ParentToClient(AObj: TGauge; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure Gauge_SendToBack(AObj: TGauge); extdecl;
begin
  AObj.SendToBack;
end;

procedure Gauge_SetBounds(AObj: TGauge; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Gauge_Show(AObj: TGauge); extdecl;
begin
  AObj.Show;
end;

procedure Gauge_Update(AObj: TGauge); extdecl;
begin
  AObj.Update;
end;

function Gauge_GetTextBuf(AObj: TGauge; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Gauge_GetTextLen(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Gauge_FindComponent(AObj: TGauge; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Gauge_GetNamePath(AObj: TGauge): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Gauge_Assign(AObj: TGauge; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Gauge_ClassName(AObj: TGauge): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Gauge_Equals(AObj: TGauge; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Gauge_GetHashCode(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Gauge_ToString(AObj: TGauge): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Gauge_GetPercentDone(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.PercentDone;
end;

function Gauge_GetAlign(AObj: TGauge): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Gauge_SetAlign(AObj: TGauge; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Gauge_GetAnchors(AObj: TGauge): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Gauge_SetAnchors(AObj: TGauge; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Gauge_GetBackColor(AObj: TGauge): TColor; extdecl;
begin
  Result :=  AObj.BackColor;
end;

procedure Gauge_SetBackColor(AObj: TGauge; AValue: TColor); extdecl;
begin
  AObj.BackColor := AValue;
end;

function Gauge_GetBorderStyle(AObj: TGauge): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Gauge_SetBorderStyle(AObj: TGauge; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function Gauge_GetColor(AObj: TGauge): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure Gauge_SetColor(AObj: TGauge; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function Gauge_GetEnabled(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Gauge_SetEnabled(AObj: TGauge; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Gauge_GetForeColor(AObj: TGauge): TColor; extdecl;
begin
  Result :=  AObj.ForeColor;
end;

procedure Gauge_SetForeColor(AObj: TGauge; AValue: TColor); extdecl;
begin
  AObj.ForeColor := AValue;
end;

function Gauge_GetFont(AObj: TGauge): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure Gauge_SetFont(AObj: TGauge; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function Gauge_GetMinValue(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.MinValue;
end;

procedure Gauge_SetMinValue(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.MinValue := AValue;
end;

function Gauge_GetMaxValue(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.MaxValue;
end;

procedure Gauge_SetMaxValue(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.MaxValue := AValue;
end;

function Gauge_GetParentColor(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure Gauge_SetParentColor(AObj: TGauge; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function Gauge_GetParentFont(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure Gauge_SetParentFont(AObj: TGauge; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function Gauge_GetParentShowHint(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Gauge_SetParentShowHint(AObj: TGauge; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function Gauge_GetPopupMenu(AObj: TGauge): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Gauge_SetPopupMenu(AObj: TGauge; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function Gauge_GetProgress(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.Progress;
end;

procedure Gauge_SetProgress(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.Progress := AValue;
end;

function Gauge_GetShowHint(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Gauge_SetShowHint(AObj: TGauge; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Gauge_GetShowText(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.ShowText;
end;

procedure Gauge_SetShowText(AObj: TGauge; AValue: LongBool); extdecl;
begin
  AObj.ShowText := AValue;
end;

function Gauge_GetVisible(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Gauge_SetVisible(AObj: TGauge; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function Gauge_GetAction(AObj: TGauge): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Gauge_SetAction(AObj: TGauge; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function Gauge_GetBiDiMode(AObj: TGauge): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Gauge_SetBiDiMode(AObj: TGauge; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure Gauge_GetBoundsRect(AObj: TGauge; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Gauge_SetBoundsRect(AObj: TGauge; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Gauge_GetClientHeight(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Gauge_SetClientHeight(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Gauge_GetClientRect(AObj: TGauge; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Gauge_GetClientWidth(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Gauge_SetClientWidth(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Gauge_GetFloating(AObj: TGauge): LongBool; extdecl;
begin
  Result :=  AObj.Floating;
end;

function Gauge_GetParent(AObj: TGauge): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Gauge_SetParent(AObj: TGauge; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Gauge_GetLeft(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Gauge_SetLeft(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Gauge_GetTop(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Gauge_SetTop(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Gauge_GetWidth(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Gauge_SetWidth(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Gauge_GetHeight(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Gauge_SetHeight(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Gauge_GetCursor(AObj: TGauge): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Gauge_SetCursor(AObj: TGauge; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Gauge_GetHint(AObj: TGauge): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Gauge_SetHint(AObj: TGauge; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Gauge_GetComponentCount(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Gauge_GetComponentIndex(AObj: TGauge): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Gauge_SetComponentIndex(AObj: TGauge; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Gauge_GetOwner(AObj: TGauge): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Gauge_GetName(AObj: TGauge): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Gauge_SetName(AObj: TGauge; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Gauge_GetTag(AObj: TGauge): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Gauge_SetTag(AObj: TGauge; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Gauge_GetComponents(AObj: TGauge; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Gauge_Create,
  Gauge_Free,
  Gauge_AddProgress,
  Gauge_BringToFront,
  Gauge_ClientToScreen,
  Gauge_ClientToParent,
  Gauge_Dragging,
  Gauge_HasParent,
  Gauge_Hide,
  Gauge_Invalidate,
  Gauge_Perform,
  Gauge_Refresh,
  Gauge_Repaint,
  Gauge_ScreenToClient,
  Gauge_ParentToClient,
  Gauge_SendToBack,
  Gauge_SetBounds,
  Gauge_Show,
  Gauge_Update,
  Gauge_GetTextBuf,
  Gauge_GetTextLen,
  Gauge_FindComponent,
  Gauge_GetNamePath,
  Gauge_Assign,
  Gauge_ClassName,
  Gauge_Equals,
  Gauge_GetHashCode,
  Gauge_ToString,
  Gauge_GetPercentDone,
  Gauge_GetAlign,
  Gauge_SetAlign,
  Gauge_GetAnchors,
  Gauge_SetAnchors,
  Gauge_GetBackColor,
  Gauge_SetBackColor,
  Gauge_GetBorderStyle,
  Gauge_SetBorderStyle,
  Gauge_GetColor,
  Gauge_SetColor,
  Gauge_GetEnabled,
  Gauge_SetEnabled,
  Gauge_GetForeColor,
  Gauge_SetForeColor,
  Gauge_GetFont,
  Gauge_SetFont,
  Gauge_GetMinValue,
  Gauge_SetMinValue,
  Gauge_GetMaxValue,
  Gauge_SetMaxValue,
  Gauge_GetParentColor,
  Gauge_SetParentColor,
  Gauge_GetParentFont,
  Gauge_SetParentFont,
  Gauge_GetParentShowHint,
  Gauge_SetParentShowHint,
  Gauge_GetPopupMenu,
  Gauge_SetPopupMenu,
  Gauge_GetProgress,
  Gauge_SetProgress,
  Gauge_GetShowHint,
  Gauge_SetShowHint,
  Gauge_GetShowText,
  Gauge_SetShowText,
  Gauge_GetVisible,
  Gauge_SetVisible,
  Gauge_GetAction,
  Gauge_SetAction,
  Gauge_GetBiDiMode,
  Gauge_SetBiDiMode,
  Gauge_GetBoundsRect,
  Gauge_SetBoundsRect,
  Gauge_GetClientHeight,
  Gauge_SetClientHeight,
  Gauge_GetClientRect,
  Gauge_GetClientWidth,
  Gauge_SetClientWidth,
  Gauge_GetFloating,
  Gauge_GetParent,
  Gauge_SetParent,
  Gauge_GetLeft,
  Gauge_SetLeft,
  Gauge_GetTop,
  Gauge_SetTop,
  Gauge_GetWidth,
  Gauge_SetWidth,
  Gauge_GetHeight,
  Gauge_SetHeight,
  Gauge_GetCursor,
  Gauge_SetCursor,
  Gauge_GetHint,
  Gauge_SetHint,
  Gauge_GetComponentCount,
  Gauge_GetComponentIndex,
  Gauge_SetComponentIndex,
  Gauge_GetOwner,
  Gauge_GetName,
  Gauge_SetName,
  Gauge_GetTag,
  Gauge_SetTag,
  Gauge_GetComponents;
