package privatestruct

type Contoh struct {
	Nama string
	Umur int
	hobi string
}

func (c Contoh) getNama() string {
	return c.Nama
}

func (c Contoh) GetNama() string {
	return c.Nama
}

func (c *Contoh) SetNama(nama string) {
	c.Nama = nama
}

func (c *Contoh) SetHobi(hobi string) {
	c.hobi = hobi
}

func (c Contoh) GetHobi() string {
	return c.hobi
}
