package exporter

type Exporter interface {
	Export(data interface{}, filename string) error
}

func GetAllExporters() []Exporter {
	return []Exporter{
		CSVExporter{},
		JsonExporter{},
	}
}