package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"time"
)

func monteCarloPiSimulation(numPoints int) (float64, plotter.XYs) {
	rand.Seed(time.Now().UnixNano())

	pointsInsideCircle := 0
	points := make(plotter.XYs, numPoints)

	for i := 0; i < numPoints; i++ {
		x := rand.Float64() * 2 - 1 // Random x coordinate in the range [-1, 1]
		y := rand.Float64() * 2 - 1 // Random y coordinate in the range [-1, 1]

		distance := math.Sqrt(x*x + y*y)

		if distance <= 1 {
			pointsInsideCircle++
		}

		points[i].X = x
		points[i].Y = y
	}

	piApproximation := 4.0 * float64(pointsInsideCircle) / float64(numPoints)
	return piApproximation, points
}

func visualisePoints(points plotter.XYs) error {
	p := plot.New()

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		return fmt.Errorf("error creating scatter plot: %w", err)
	}

	p.Add(scatter)
	// Creates a scatter plot for points inside the circle (use different color)
	var insideCirclePoints plotter.XYs
	for _, point := range points {
		if point.X*point.X+point.Y*point.Y <= 1 {
			insideCirclePoints = append(insideCirclePoints, point)
		}
	}
	insideCircleScatter, err := plotter.NewScatter(insideCirclePoints)
	if err != nil {
		return fmt.Errorf("error creating scatter plot for inside circle: %w", err)
	}
	insideCircleScatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255} // Sets colour 
	p.Add(insideCircleScatter)


	p.Title.Text = "Scatter Plot"
	p.X.Label.Text = "X-axis"
	p.Y.Label.Text = "Y-axis"



	// Saves the image in the same directory as the executable

	filePath := "montecarlo_scatter_plot.png"


	
	fmt.Printf("Saving Visualization to: %s\n", filePath)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, filePath); err != nil {
		return fmt.Errorf("error saving visualization: %w", err)
	}

	fmt.Printf("Visualization saved as %s\n", filePath)
	return nil
}
func main() {
	numPoints := 1000
	estimatedPi, points := monteCarloPiSimulation(numPoints)

	fmt.Printf("Estimated value of π using %d points: %f\n", numPoints, estimatedPi)

	if err := visualisePoints(points); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
