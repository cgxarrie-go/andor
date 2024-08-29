package elementtype

type Type int16

const (
	Unknown Type = 0
	Item    Type = 1
	And     Type = 2
	Or      Type = 3
)

var enumMapName = map[Type]string{
	And:  `and`,
	Or:   `or`,
	Item: `item`,
}

// Name return the name of the AOType
func (t Type) Name() string {
	return enumMapName[t]
}

// FromName returns a enumMapName matching the provided name.
func FromName(name string) Type {
	for id, n := range enumMapName {
		if n == name {
			return id
		}
	}

	return Unknown
}
