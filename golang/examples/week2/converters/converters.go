package converters

const (
  mitoKM = 1.6
  fttoM = 3.28
)

func MilesToKM(miles float64) float64 {
  return miles*mitoKM
}

func FeetToMeters(feet float64) float64 {
  return feet*fttoM
}
