package pagescale

import (
	"fmt"
	"testing"

	"github.com/timdrysdale/geo"
)

type isDynamic bool

var (
	a4pImagePath = "./img/a4p.jpg"
	a4lImagePath = "./img/a4l.jpg"
	sImagePath   = "./img/s.jpg"
	a4pDim       = geo.Dim{W: 210 * geo.PPMM, H: 297 * geo.PPMM}
	a4lDim       = geo.Dim{W: 297 * geo.PPMM, H: 210 * geo.PPMM}
	sDim         = geo.Dim{W: 150 * geo.PPMM, H: 150 * geo.PPMM}
	pageSizes    = []geo.Dim{a4pDim, a4lDim, sDim}
	imagePaths   = []string{a4pImagePath, a4lImagePath, sImagePath}
	pageTypes    = []isDynamic{true, false}
	pageNames    = []string{"page-a4p", "page-a4l", "page-s"}
	imgNames     = []string{"image-a4p", "image-a4l", "image-s"}
	typeNames    = []string{"dynamic", "static"}
)

func makePath(image string, pagetype string, pagesize string) string {
	return fmt.Sprintf("%s-%s-%s.pdf", image, pagetype, pagesize)
}

func TestPrintImageOnStaticPage(t *testing.T) {
	// These are visual tests for a human to look at....
	for imgidx, img := range imagePaths {
		for ptidx, pt := range pageTypes {
			for psidx, ps := range pageSizes {

				pdfPath := makePath(imgNames[imgidx], typeNames[ptidx], pageNames[psidx])

				if pt {
					err := PrintImageOnDynamicPage(img, ps, pdfPath)
					if err != nil {
						t.Error(err)
					}
					fmt.Printf("Dynamic page printed to %s\n", pdfPath)
				} else {
					err := PrintImageOnStaticPage(img, ps, pdfPath)
					if err != nil {
						t.Error(err)
					}
					fmt.Printf("Static page printed to %s\n", pdfPath)

				}
			}

		}

	}
}
