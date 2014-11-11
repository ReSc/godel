package config

const (
	Bootstrap string = `<?xml version="1.0" encoding="UTF-8"?>
<config name="bootstrap">
	<node name="system" >
		<node name="type">
		    <node name="none"      type="primitive" />
		    <node name="any"       type="primitive" />
		    <node name="some"      type="primitive" />
		    <node name="unit"      type="primitive" />
		    <node name="bool"      type="primitive" />
			<node name="int8"      type="primitive" />
			<node name="int16"     type="primitive" />
			<node name="int32"     type="primitive" />
			<node name="int64"     type="primitive" />
			<node name="uint8"     type="primitive" />
			<node name="uint16"    type="primitive" />
			<node name="uint32"    type="primitive" />
			<node name="uint64"    type="primitive" />
			<node name="float32"   type="primitive" />
			<node name="float64"   type="primitive" />
			<node name="decimal"   type="primitive" />
			<node name="uuid"      type="primitive" />
			<node name="string"    type="primitive" />
		</node>
	</node> 
</config>`
)
