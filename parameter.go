package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include <freetype/ftparams.h>
//
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type Parameter = C.FT_Parameter

func (param Parameter) freeData() {
	C.free(unsafe.Pointer(param.data))
}

func ParamTagIgnoreTypoGraphicFamily(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_FAMILY, value)
}

func ParamTagIgnoreTypoGraphicSubfamily(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_SUBFAMILY, value)
}

func ParamTagIncremental(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_INCREMENTAL, value)
}

func ParamTagIgnoreSbix(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_SBIX, value)
}

const lcdFilterWeightsLen = 5

func ParamTagLCDFilterWeights(weights *[lcdFilterWeightsLen]byte) Parameter {
	if weights == nil {
		return C.FT_Parameter{
			tag:  C.FT_PARAM_TAG_LCD_FILTER_WEIGHTS,
			data: nil,
		}
	}

	cWeights := (*C.uchar)(C.malloc(lcdFilterWeightsLen))
	C.memcpy(unsafe.Pointer(cWeights), unsafe.Pointer(&(*weights)[0]), lcdFilterWeightsLen)

	return C.FT_Parameter{
		tag:  C.FT_PARAM_TAG_LCD_FILTER_WEIGHTS,
		data: C.FT_Pointer(cWeights),
	}
}

func ParamTagRandomSeed(value *int) Parameter {
	return integerParamTag(C.FT_PARAM_TAG_RANDOM_SEED, value)
}

func ParamTagStemDarkening(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_STEM_DARKENING, value)
}

func ParamTagUnpatentedHinting(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_UNPATENTED_HINTING, value)
}

func booleanParamTag(tag C.FT_ULong, value *bool) Parameter {
	if value == nil {
		return C.FT_Parameter{
			tag:  tag,
			data: nil,
		}
	}

	var cBool C.FT_Bool
	cValue := (*C.FT_Bool)(C.malloc(C.size_t(unsafe.Sizeof(cBool))))
	if *value {
		*cValue = 1
	} else {
		*cValue = 0
	}

	return C.FT_Parameter{
		tag:  tag,
		data: C.FT_Pointer(cValue),
	}
}

func integerParamTag(tag C.FT_ULong, value *int) Parameter {
	if value == nil {
		return C.FT_Parameter{
			tag:  tag,
			data: nil,
		}
	}

	var cInt C.FT_Int32
	cValue := (*C.FT_Int32)(C.malloc(C.size_t(unsafe.Sizeof(cInt))))
	*cValue = C.FT_Int32(*value)

	return C.FT_Parameter{
		tag:  tag,
		data: C.FT_Pointer(cValue),
	}
}
