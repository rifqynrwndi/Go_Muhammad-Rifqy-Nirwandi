package services

type MathService struct{}

func NewMathService() *MathService {
    return &MathService{}
}

func (s *MathService) Add(x, y float64) float64 {
    return x + y
}

func (s *MathService) Subtract(x, y float64) float64 {
    return x - y
}

func (s *MathService) Multiply(x, y float64) float64 { // Ensure Multiply is defined
    return x * y
}

func (s *MathService) Divide(x, y float64) float64 {
    if y == 0 {
        return 0 // Handle division by zero
    }
    return x / y
}
