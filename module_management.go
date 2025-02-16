package freetype

import (
	"fmt"

	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// How to add, upgrade, remove, and control modules from FreeType.

// FT_Module

// FT_Module_Constructor

// FT_Module_Destructor

// FT_Module_Requester

// FT_Module_Class

// FT_Add_Module

// FT_Get_Module

// FT_Remove_Module

// FT_Add_Default_Modules

// FT_FACE_DRIVER_NAME

// Property_Sets set a property for a given module.
//
// https://freetype.org/freetype2/docs/reference/ft2-module_management.html#ft_property_set
func (lib Library) PropertySet(moduleName string, propertyName string, value uintptr) error {
	cModuleName, err := libc.CString(moduleName)
	if err != nil {
		return fmt.Errorf("failed to set create C string for module name %s : %w", moduleName, err)
	}
	defer libc.Xfree(nil, cModuleName)

	cPropertyName, err := libc.CString(propertyName)
	if err != nil {
		return fmt.Errorf("failed to set create C string for property name %s : %w", propertyName, err)
	}
	defer libc.Xfree(nil, cPropertyName)

	err_ := libfreetype.XFT_Property_Set(lib.tls, lib.library, cModuleName, cPropertyName, value)
	return newError(err_, "failed to set property %s for module %s", propertyName, moduleName)
}

// FT_Property_Get

// FT_Set_Default_Properties

// FT_New_Library

// FT_Done_Library

// FT_Reference_Library

// FT_Renderer

// FT_Renderer_Class

// FT_Get_Renderer

// FT_Set_Renderer

// FT_Set_Debug_Hook

// FT_Driver

// FT_DebugHook_Func

// FT_DEBUG_HOOK_XXX
