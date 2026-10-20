package spinner

var sizeClasses = map[Size]string{
	SizeSM: "size-4",
	SizeMD: "size-5 sm:size-4",
	SizeLG: "size-8",
}

// svgClasses spins the SVG. Under prefers-reduced-motion it keeps
// turning, slowly, rather than stopping — a frozen spinner reads as
// "stuck", which is worse than gentle motion.
func svgClasses(s Size, extra string) string {
	size, ok := sizeClasses[s]
	if !ok {
		size = sizeClasses[SizeMD]
	}
	return join("shrink-0 animate-spin motion-reduce:[animation-duration:3s] "+size, extra)
}

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
