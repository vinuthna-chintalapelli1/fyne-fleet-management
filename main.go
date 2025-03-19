package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fleet Management")

	// Define the label here so it's accessible in the callback
	selectedLabel := widget.NewLabel("No region selected")

	// Create a dropdown for selecting regions
	regions := []string{"US", "EU", "JP"}
	regionSelector := widget.NewSelect(regions, func(selected string) {
		// Update the label text when a region is selected
		selectedLabel.SetText("Region selected: " + selected)
	})

	// Create table data for regions
	customers := []struct {
		Name   string
		Region string
	}{
		{"Acme", "US"},
		{"Beta", "EU"},
		{"Gamma", "JP"},
	}

	// Create a table to display customer data
	table := widget.NewTable(
		func() (int, int) { return len(customers), 2 }, // 2 columns (Name and Region)
		func() fyne.CanvasObject {
			// Each row will have two labels, one for Name and one for Region
			nameLabel := widget.NewLabel("")
			regionLabel := widget.NewLabel("")
			// Set the minimum height for the rows by using container.NewVBox
			rowContainer := container.NewVBox(nameLabel, regionLabel)
			rowContainer.Resize(fyne.NewSize(0, 80)) // Adjust the row height to 80
			return rowContainer
		},
		func(cell widget.TableCellID, o fyne.CanvasObject) {
			// Populate the table with customer data
			customer := customers[cell.Row]
			switch cell.Col {
			case 0:
				o.(*fyne.Container).Objects[0].(*widget.Label).SetText(customer.Name)
			case 1:
				o.(*fyne.Container).Objects[1].(*widget.Label).SetText(customer.Region)
			}
		},
	)

	// Add padding to the top panel for more height
	topPanel := container.NewVBox(
		regionSelector,
		selectedLabel,
	)

	// Add components to the window
	myWindow.SetContent(container.NewVBox(topPanel, table))
	myWindow.Resize(fyne.NewSize(800, 600)) // Resize the window to make sure it's big enough for the table

	// Show the window
	myWindow.ShowAndRun()
}
