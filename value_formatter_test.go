package chart

import (
	"testing"
	"time"

	"github.com/StephanVerbeeck/go-chart/v2/testutil"
)

func TestTimeValueFormatterWithFormat(t *testing.T) {
	// replaced new assertions helper

	d := time.Now()
	di := TimeToFloat64(d)
	df := float64(di)

	s := FormatTime(d, DefaultDateFormat)
	si := FormatTime(di, DefaultDateFormat)
	sf := FormatTime(df, DefaultDateFormat)
	testutil.AssertEqual(t, s, si)
	testutil.AssertEqual(t, s, sf)

	sd := DateValueFormatter(d)
	sdi := DateValueFormatter(di)
	sdf := DateValueFormatter(df)
	testutil.AssertEqual(t, s, sd)
	testutil.AssertEqual(t, s, sdi)
	testutil.AssertEqual(t, s, sdf)
}

func TestFloatValueFormatter(t *testing.T) {
	// replaced new assertions helper
	testutil.AssertEqual(t, "1234.00", FloatValueFormatter(1234.00))
}

func TestFloatValueFormatterWithFloat32Input(t *testing.T) {
	// replaced new assertions helper
	testutil.AssertEqual(t, "1234.00", FloatValueFormatter(float32(1234.00)))
}

func TestFloatValueFormatterWithIntegerInput(t *testing.T) {
	// replaced new assertions helper
	testutil.AssertEqual(t, "1234.00", FloatValueFormatter(1234))
}

func TestFloatValueFormatterWithInt64Input(t *testing.T) {
	// replaced new assertions helper
	testutil.AssertEqual(t, "1234.00", FloatValueFormatter(int64(1234)))
}

func TestFloatValueFormatterWithFormat(t *testing.T) {
	// replaced new assertions helper

	v := 123.456
	sv := FloatValueFormatterWithFormat(v, "%.3f")
	testutil.AssertEqual(t, "123.456", sv)
	testutil.AssertEqual(t, "123.000", FloatValueFormatterWithFormat(123, "%.3f"))
}
