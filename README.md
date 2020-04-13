# pagescale
scaling formulas for fitting images to static and dynamic gradex pages

## static vs dynamic page width

pdf.gradex™ has both static and dynamic page widths. The height is always fixed - one height for portrait, the other for landscape. 
It's not a fundamental limitation, but for convenience, a single paper size is assumed for a given exam. It's electronic marking, at least partly, so you can always zoom in if someone sends you A3 with teeny tiny writing and it is scaled to A4. So then we get a mix of page sizes coming in - which we want to scale to page height for a dynamic page, or until the page hits the width limit of a static page.No rocket science there - just wanted a lightweight repo in which to confirm it works before embedding into the bigger work.

## test output

![alt text][montage]

[montage]: ./img/montage.png "montage of 18 images imported into 18 pages"
