package stuartclient

type StuartPackageDimensions struct {
	size   PackageType
	volume float32
	weight float32
}

var stuartPackageDefinitions = []StuartPackageDimensions{
	{
		size:   Small,
		volume: 40.0 * 20.0 * 15.0,
		weight: 12.0,
	},
	{
		size:   Medium,
		volume: 50.0 * 30.0 * 30.0,
		weight: 12.0,
	},
	{
		size:   Large,
		volume: 90.0 * 65.0 * 50.0,
		weight: 25.0,
	},
	{
		size:   Xlarge,
		volume: 100.0 * 90.0 * 50.0,
		weight: 70.0,
	},
}

func CalculateParcelSize(heightCM float32, lengthCm float32, widthCm float32, weightKg float32) PackageType {
	volume := heightCM * lengthCm * widthCm
	for _, definition := range stuartPackageDefinitions {
		if definition.volume >= volume && definition.weight >= weightKg {
			return definition.size
		}
	}
	//Didn't find any. Go for extra-large to be sure
	return Xlarge
}
