package main

import (
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// This function will start the toolbox of URLs
func openToolbox() fyne.CanvasObject {
	// Create URL objects for different services
	confluenceURL, err := url.Parse("https://confluence.storage.hpecorp.net/display/NCS/DSCC")

	if err != nil {
		log.Println(err)
		return nil
	}

	grafanaURL, err := url.Parse("https://fleetpoc2-us-west-2.cloudops.qa.cds.hpe.com/grafana/d/uid_search_data_processor/search-data-processor?orgId=1")

	if err != nil {
		log.Println(err)
		return nil
	}

	humioURL, err := url.Parse("https://fleetpoc2-us-west-2.cloudops.qa.cds.hpe.com/logs/storagecentral/search")

	if err != nil {
		log.Println(err)
		return nil
	}

	pavoURL, err := url.Parse("https://console-neonops3-app.qa.cds.hpe.com/")

	if err != nil {
		log.Println(err)
		return nil
	}

	// Create content for the toolbox
	hello := widget.NewLabel("Welcome to the toolbox!")

	// Create the toolbox content
	toolboxContent := container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome to the handy links!")
		}),
		widget.NewHyperlink("Confluence", confluenceURL),
		widget.NewHyperlink("Grafana", grafanaURL),
		widget.NewHyperlink("Humio", humioURL),
		widget.NewHyperlink("Pavo / Aquilla", pavoURL),
	)

	return toolboxContent
}

func main() {
	// Initialize Fyne app and main window
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

	// Create a tab container
	tabs := container.NewAppTabs(
		container.NewTabItem("Fleet Management", container.NewVBox(regionSelector, selectedLabel, table)),
		container.NewTabItem("Toolbox", openToolbox()), // Open Toolbox in the second tab
	)

	// Set the content of the window
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(800, 600)) // Resize the window to make sure it's big enough for the table

	// Show the window
	myWindow.ShowAndRun()
}
