package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	SHAPE_TYPE_NULL        = 0
	SHAPE_TYPE_POINT       = 1
	SHAPE_TYPE_POLYLINE    = 3
	SHAPE_TYPE_POLYGON     = 5
	SHAPE_TYPE_MULTIPOINT  = 8
	SHAPE_TYPE_POINTZ      = 11
	SHAPE_TYPE_POLYLINEZ   = 13
	SHAPE_TYPE_POLYGONZ    = 15
	SHAPE_TYPE_MULTIPOINTZ = 18
	SHAPE_TYPE_POINTM      = 21
	SHAPE_TYPE_POLYLINEM   = 23
	SHAPE_TYPE_POLYGONM    = 25
	SHAPE_TYPE_MULTIPOINTM = 28
	SHAPE_TYPE_MULTIPATCH  = 31
)

type BigHeader struct {
	MagicNumber int32
	Zeros       [20]byte
	FileLen     int32
}

type LittleHeader struct {
	Version   int32
	ShapeType int32
	XMin      float64
	XMax      float64
	YMin      float64
	YMax      float64
	ZMin      float64
	ZMax      float64
	MMin      float64
	MMax      float64
}

func detectShapeType(shapeType int32) string {
	switch shapeType {
	case SHAPE_TYPE_NULL:
		return "NULL"
	case SHAPE_TYPE_POINT:
		return "POINT"
	case SHAPE_TYPE_POLYLINE:
		return "POLYLINE"
	case SHAPE_TYPE_POLYGON:
		return "POLYGON"
	case SHAPE_TYPE_MULTIPOINT:
		return "MULTIPOINT"
	case SHAPE_TYPE_POINTZ:
		return "POINTZ"
	case SHAPE_TYPE_POLYLINEZ:
		return "POLYLINEZ"
	case SHAPE_TYPE_POLYGONZ:
		return "POLYGONZ"
	case SHAPE_TYPE_MULTIPOINTZ:
		return "MULTIPOINTZ"
	case SHAPE_TYPE_POINTM:
		return "POINTM"
	case SHAPE_TYPE_POLYLINEM:
		return "POLYLINEM"
	case SHAPE_TYPE_POLYGONM:
		return "POLYGONM"
	case SHAPE_TYPE_MULTIPOINTM:
		return "MULTIPOINTM"
	case SHAPE_TYPE_MULTIPATCH:
		return "MULTIPATCH"
	default:
		return "UNKNOWN"
	}
}

func main() {
	if 2 < len(os.Args) {
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()
	firstHeader := new(BigHeader)
	secondHeader := new(LittleHeader)
	binary.Read(f, binary.BigEndian, firstHeader)
	binary.Read(f, binary.LittleEndian, secondHeader)
	if firstHeader.MagicNumber != 9994 && secondHeader.Version != 1000 {
		fmt.Println("invalid")
	}
	fmt.Printf("file size: %dkb\n", firstHeader.FileLen/1024)
	fmt.Printf("features type: %d(%s)\n", secondHeader.ShapeType, detectShapeType(secondHeader.ShapeType))
	fmt.Printf("bbox(x axis): %f～%f\n", secondHeader.XMin, secondHeader.XMax)
	fmt.Printf("bbox(y axis): %f～%f\n", secondHeader.YMin, secondHeader.YMax)
	fmt.Printf("bbox(z axis): %f～%f\n", secondHeader.ZMin, secondHeader.ZMax)
	fmt.Printf("bbox(m): %f～%f\n", secondHeader.MMin, secondHeader.MMax)
}
