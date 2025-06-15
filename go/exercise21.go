package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" {
		return false
	}
	if r.Address == nil {
		return false
	}
	street, hasStreet := r.Address["street"]
	return hasStreet && street != ""
}

func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

func Count(residents []*Resident) int {
	count := 0
	for _, resident := range residents {
		if resident != nil && resident.HasRequiredInfo() {
			count++
		}
	}
	return count
}
