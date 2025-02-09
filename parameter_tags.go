package freetype

import "modernc.org/libfreetype"

// ParamTag is a tag for use with the Parameter type.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#parameter-tags
type ParamTag = libfreetype.TFT_ULong

var (
	PARAM_TAG_IGNORE_TYPOGRAPHIC_FAMILY    = ParamTag(imageTag('i', 'g', 'p', 'f'))
	PARAM_TAG_IGNORE_TYPOGRAPHIC_SUBFAMILY = ParamTag(imageTag('i', 'g', 'p', 's'))
	PARAM_TAG_INCREMENTAL                  = ParamTag(imageTag('i', 'n', 'c', 'r'))
	PARAM_TAG_IGNORE_SBIX                  = ParamTag(imageTag('i', 's', 'b', 'x'))
	PARAM_TAG_LCD_FILTER_WEIGHTS           = ParamTag(imageTag('l', 'c', 'd', 'f'))
	PARAM_TAG_RANDOM_SEED                  = ParamTag(imageTag('s', 'e', 'e', 'd'))
	PARAM_TAG_STEM_DARKENING               = ParamTag(imageTag('d', 'a', 'r', 'k'))
	PARAM_TAG_UNPATENTED_HINTING           = ParamTag(imageTag('u', 'n', 'p', 'a'))
)
