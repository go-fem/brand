# go-fem — brand assets

Official logos for the **go-fem** organization.

<img src="avatar/go-fem.png" width="160" alt="go-fem">

The mark is a mesh: a square of 3 x 3 nodes, every cell cut into two triangles, the elements a PDE is solved on. It is written by `mark/main.go`, which lays the points out from a short
table; nothing is placed by hand. Its gradient (`#75DC6C` to `#15820B`) is a
green (hue 115) that sets go-fem apart from the other go-scikit organizations,
including the blue-green of go-control (142) and the olive of go-autodiff (87);
its dark end reads on white at 4.96:1 and its light end on a dark page at
11.25:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-fem.svg` differs.

---
*Generated from `svg/color/go-fem.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
