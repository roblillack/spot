package ui

import "github.com/roblillack/spot"

func calcLayout(parent spot.Control, cX, cY, cW, cH int) (int, int, int, int) {
	if parent == nil {
		return cX, cY, cW, cH
	}

	if cX >= 0 && cY >= 0 && cW > 0 && cH > 0 {
		return cX, cY, cW, cH
	}

	container, ok := parent.(spot.Container)
	if !ok || container == nil {
		return cX, cY, cW, cH
	}

	var x, y, w, h int
	pW := container.ContentWidth()
	pH := container.ContentHeight()

	// Horizontal layout
	if cX >= 0 && cW == 0 {
		// Left anchored, fill width
		x = cX
		w = pW - cX
	} else if cX < 0 && cW == 0 {
		// Right anchored, fill width
		x = 0
		w = pW + cX
	} else if cW < 0 {
		// Left & Right anchored, filled width
		if cX < 0 {
			cX = -cX
		}
		x = cX
		w = pW + cW - cX
	} else if cX < 0 && cW > 0 {
		// Right anchored, fixed width
		x = pW + cX - cW
		w = cW
	} else if cX >= 0 && cW > 0 {
		// Left anchored, fixed width
		x = cX
		w = cW
	}

	// Vertical layout
	if cY >= 0 && cH == 0 {
		// Top anchored, fill height
		y = cY
		h = pH - cY
	} else if cY < 0 && cH == 0 {
		// Bottom anchored, fill height
		y = 0
		h = pH + cY
	} else if cH < 0 {
		// Top & Bottom anchored, filled height
		if cY < 0 {
			cY = -cY
		}
		y = cY
		h = pH + cH - cY
	} else if cY < 0 && cH > 0 {
		// Bottom anchored, fixed height
		y = pH + cY - cH
		h = cH
	} else if cY >= 0 && cH > 0 {
		// Top anchored, fixed height
		y = cY
		h = cH
	}

	return x, y, w, h
}
