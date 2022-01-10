package onnxReference

import (
	"backend/images"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gorgonia.org/tensor"
)

func OnnxRef(inpath string, outpath string) {
	// I. Inputの作成
	file, err := os.Open(inpath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	img, format, err := image.Decode(file)
	log.Println("ImageFormat:", format)
	if err != nil {
		log.Fatal(err)
	}

	height, width := img.Bounds().Dy(), img.Bounds().Dx()
	input := tensor.New(tensor.WithShape(1, 3, height, width), tensor.Of(tensor.Float32))
	err = images.ImageToBCHW(img, input)
	if err != nil {
		log.Fatal(err)
	}

	// II. Modelの作成
	backend := gorgonnx.NewGraph()
	model := onnx.NewModel(backend)

	byte_model, _ := os.ReadFile("./SampleModel.onnx")
	err = model.UnmarshalBinary(byte_model)
	if err != nil {
		log.Fatal(err)
	}

	// III. Inference
	model.SetInput(0, input)
	err = backend.Run()
	if err != nil {
		log.Fatal(err)
	}
	output, _ := model.GetOutputTensors()

	outimg, err := images.TensorToImg(output[0])
	if err != nil {
		log.Fatal(err)
	}

	// IV. 予測の保存
	file_out, err := os.Create(outpath)
	defer file_out.Close()
	if err != nil {
		log.Fatal(err)
	}
	if format == "png" {
		png.Encode(file_out, outimg)
	} else {
		jpeg.Encode(file_out, outimg, nil)
	}
}
