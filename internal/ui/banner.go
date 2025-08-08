/*
©cyber-atharv | 2026
banner.go

ASCII art banner renderer for the sentinel CLI header

Renders the SENTINEL logo in alternating red and cyan before scans.
*/

package ui

import "fmt"

var sentinelBanner = []string{
	"███████╗███████╗███╗   ██╗████████╗██╗███╗   ██╗███████╗██╗",
	"██╔════╝██╔════╝████╗  ██║╚══██╔══╝██║████╗  ██║██╔════╝██║",
	"███████╗█████╗  ██╔██╗ ██║   ██║   ██║██╔██╗ ██║█████╗  ██║",
	"╚════██║██╔══╝  ██║╚██╗██║   ██║   ██║██║╚██╗██║██╔══╝  ██║",
	"███████║███████╗██║ ╚████║   ██║   ██║██║ ╚████║███████╗███████╗",
	"╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝╚═╝  ╚═══╝╚══════╝╚══════╝",
}

var bannerColors = []func(a ...any) string{
	Cyan,
	Red,
	Cyan,
	Red,
	Cyan,
	Red,
}

func PrintBanner() {
	fmt.Println()
	for i, line := range sentinelBanner {
		c := bannerColors[i%len(bannerColors)]
		fmt.Printf("  %s\n", c(line))
	}
	fmt.Printf(
		"  %s\n",
		HiBlackItalic(
			"  Linux persistence mechanism scanner",
		),
	)
	fmt.Printf("  %s\n\n", Dim(HRule(65)))
}
