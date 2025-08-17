package cli

import "fmt"

func printLogo() {
	fmt.Println(pink("    =/\\                 /\\="))
	fmt.Println(pink("    / \\'._   (\\_/)   _.'/ \\"))
	fmt.Printf("   %s%s%s%s%s          %s %s\n",
		pink("/ .''._'--("),
		white("o"),
		pink("."),
		white("o"),
		pink(")--'_.''. \\"),
		white("BatCrack"),
		pink("v1.0.0"),
	)
	fmt.Printf("  %s         %s %s\n",
		pink("/.' _/ |`'=/ \" \\='`| \\_ `.\\"),
		white("(c) 2025"),
		pink("MomboteQ"),
	)
	fmt.Printf(" %s        %s\n",
		pink("/` .' `\\;-,'\\___/',-;/` '. '\\"),
		white("MIT License"),
	)
	fmt.Println(pink("/.-' jgs   `\\(-V-)/`       `-.\\"))
	fmt.Println(pink("`            \"   \"            `"))
	fmt.Println()
}
