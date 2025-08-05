package main

import (
	"fmt"

	"github.com/muesli/termenv"
)

func infoTitle(msg string) string {
	return termenv.String(msg).Foreground(termenv.ColorProfile().Color("2")).Bold().String()
}

func warnTitle(msg string) string {
	return termenv.String(msg).Foreground(termenv.ColorProfile().Color("3")).Bold().String()
}

func errorTitle(msg string) string {
	return termenv.String(msg).Foreground(termenv.ColorProfile().Color("9")).Bold().String()
}

func infoText(msg string) string {
	return termenv.String(msg).Foreground(termenv.ColorProfile().Color("2")).String()
}

func warnText(msg string) string {
	return termenv.String(msg).Foreground(termenv.ColorProfile().Color("3")).String()
}

func errorText(msg string) string {
	return termenv.String(msg).Foreground(termenv.ColorProfile().Color("9")).String()
}

func print(msg string) {
	fmt.Printf("\n  " + msg + "\n\n")
}
