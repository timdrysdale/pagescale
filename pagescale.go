package pagescale

import (
	"errors"
	"fmt"

	"github.com/timdrysdale/geo"
	"github.com/unidoc/unipdf/v3/creator"
)

func PrintImageOnDynamicPage(imgPath string, pageDim geo.Dim, pdfPath string) error {

	c := creator.New()

	img, err := c.NewImageFromFile(imgPath)

	if err != nil {
		return errors.New(fmt.Sprintf("Error opening image file: %s", err))
	}

	img.ScaleToHeight(pageDim.H)
	pageDim.W = img.Width()
	img.SetPos(0, 0) //top left 	c.SetPageMargins(0, 0, 0, 0) // we're not printing, so margin needed
	c.SetPageSize(creator.PageSize{pageDim.W, pageDim.H})
	c.NewPage()
	c.Draw(img)
	c.WriteToFile(pdfPath)

	return nil
}

func PrintImageOnStaticPage(imgPath string, pageDim geo.Dim, pdfPath string) error {

	c := creator.New()

	img, err := c.NewImageFromFile(imgPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error opening image file: %s", err))
	}

	// estimate image width if we scale to full height
	// note to future self - don't use the library to handle
	// an actual scaling operation as a means to do the calc
	// because if you have to rescale, there may be unintended
	// effects due to changes under the hood of the library
	// (it'd be ok with current implementation, but ...)
	imgScaledWidth := img.Width() * pageDim.H / img.Height()

	if imgScaledWidth > pageDim.W {
		// oops, we're too big, so scale using width instead
		img.ScaleToWidth(pageDim.W)
	} else {
		img.ScaleToHeight(pageDim.H)
	}

	img.SetPos(0, 0) //top left

	c.SetPageMargins(0, 0, 0, 0) // we're not printing, so margin needed
	c.SetPageSize(creator.PageSize{pageDim.W, pageDim.H})
	c.NewPage()
	c.Draw(img)
	c.WriteToFile(pdfPath)

	return nil
}
