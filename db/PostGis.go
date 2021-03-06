package db

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/linqiurong2021/go-postgis/conf"
	"github.com/linqiurong2021/go-postgis/libs/consts"
)

// PostGis PostGis
type PostGis struct {
	Pool *pgxpool.Pool
}

// NewPool 建立数据库连接
func NewPool() *pgxpool.Pool {
	fmt.Printf("%v", conf.Conf)
	postgre := NewPostgreDB(conf.Conf)
	dbpool, err := postgre.Connect()
	if err != nil {
		fmt.Printf("connect postgre err:%s", err)
		os.Exit(1)
	}
	return dbpool
}

// NewPostGis NewPostGis
func NewPostGis() *PostGis {
	return &PostGis{
		Pool: NewPool(),
	}
}

// getReturnType 获取调用返回的类型
func (r *PostGis) getReturnType(returnType string, SQL string) string {
	/*
		ST_AsBinary(geometry,{'NDR'|'XDR'})
		ST_AsEWKT(geometry)
		ST_AsEWKB(geometry, {'NDR'|'XDR'})
		ST_AsHEXEWKB(geometry, {'NDR'|'XDR'})
		ST_AsSVG(geometry, [rel], [precision])
		ST_AsGML([version], geometry, [precision])
		ST_AsKML([version], geometry, [precision])
		ST_AsGeoJson([version], geometry, [precision], [options])
	*/
	// 转小写
	returnType = strings.ToLower(returnType)
	// 函数字符串
	funcStr := ""
	// 判断
	switch returnType {
	case consts.KML:
		funcStr = fmt.Sprintf("ST_AsKML(%s)", SQL)
	case consts.GeoJSON:
		funcStr = fmt.Sprintf("ST_AsGeoJson(%s)", SQL)
	case consts.HEXEWKB:
		funcStr = fmt.Sprintf("ST_AsHEXEWKB(%s)", SQL)
	case consts.SVG:
		funcStr = fmt.Sprintf("ST_AsSVG(%s)", SQL)
	case consts.Binary:
		funcStr = fmt.Sprintf("ST_AsBinary(%s)", SQL)
	case consts.GML:
		funcStr = fmt.Sprintf("ST_AsGML(%s)", SQL)
	default:
		funcStr = fmt.Sprintf("ST_AsText(%s)", SQL)
	}
	// 返回数据
	return funcStr
}

// getCoordinate getCoordinate
func (r *PostGis) stTransform(geometry interface{}, coordinateType string, from int64, to int64) (tmpStr string) {
	//
	if coordinateType == consts.Geographic {
		tmpStr = fmt.Sprintf("st_transform(ST_SetSRID(ST_AsText('%s'),%d), %d)", geometry, from, to)
	} else {
		tmpStr = fmt.Sprintf("ST_AsText('%s')", geometry)
	}

	return tmpStr
}

// getGeomCoord getGeomCoord
func (r *PostGis) getGeomCoord(geometry interface{}, coordinateType string, from int64, to int64) interface{} {
	//
	if from == 0 {
		from = 4326 // 84坐标
	}
	if to == 0 {
		to = 4527 // 2000坐标
	}
	// 函数字符串
	funcStr := r.stTransform(geometry, coordinateType, from, to)
	// 判断
	return funcStr
}

// Centroid Centroid
func (r *PostGis) Centroid(geometry interface{}, returnType string) (result interface{}, err error) {
	// 中心坐标点
	tmpStr := fmt.Sprintf("ST_Centroid('%s')", geometry)
	// 获取返回格式
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s as Center;", SQL)
	if err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result); err != nil {
		return nil, err
	}
	return
}

// Area 面积(单位:平方米)
func (r *PostGis) Area(geometry interface{}, coordinateType string, from int64, to int64) (result interface{}, err error) {
	// 判断是否是地理坐标系 还是 投影坐标系统
	coordinateType = strings.ToLower(coordinateType)
	// 如果是地理坐标系统需要转换成投影坐标系统
	tmpStr := r.getGeomCoord(geometry, coordinateType, from, to)

	SQL := fmt.Sprintf("SELECT ST_Area(%s);", tmpStr)

	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	return
}

// Length 长度(单位:米)
func (r *PostGis) Length(geometry interface{}, coordinateType string, from int64, to int64) (result interface{}, err error) {
	//
	coordinateType = strings.ToLower(coordinateType)
	// 如果是地理坐标系统需要转换成投影坐标系统
	tmpStr := r.getGeomCoord(geometry, coordinateType, from, to)
	SQL := fmt.Sprintf("SELECT ST_Length(%s);", tmpStr)

	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// PointOnSurface PointOnSurface
func (r *PostGis) PointOnSurface(geometry interface{}, returnType string) (result interface{}, err error) {
	//
	// 获取返回格式
	tmpStr := fmt.Sprintf("ST_PointOnSurface('%s')", geometry)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	// fmt.Printf("SQL:%s", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	return
}

// Boundary 边界
func (r *PostGis) Boundary(geometry interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_Boundary('%s')", geometry)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	// fmt.Printf("SQL:%s", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// Buffer 缓冲
func (r *PostGis) Buffer(geometry interface{}, returnType string, distance float64) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_Buffer('%s',%f)", geometry, distance)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// ConvexHull 返回几何值的凸包
func (r *PostGis) ConvexHull(geometry interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_ConvexHull('%s')", geometry)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// Intersection 交点
func (r *PostGis) Intersection(geometry interface{}, geometry2 interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_Intersection('%s','%s')", geometry, geometry2)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// ShiftLongitude 读取几何图形中每个要素的每个组件中的每个点/顶点，如果经度坐标<0，则将其添加360。结果将是要在180中心图中绘制的数据的0-360版本
func (r *PostGis) ShiftLongitude(geometry interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_ShiftLongitude('%s')", geometry)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// SymDifference 返回表示A和B不相交的部分的几何。之所以称为对称差异，
// 是因为ST_SymDifference（A，B）= ST_SymDifference（B，A）。可以将其视为ST_Union（geomA，geomB）-ST_Intersection（A，B）。
func (r *PostGis) SymDifference(geometry interface{}, geometry2 interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_SymDifference('%s','%s')", geometry, geometry2)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)

	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// Difference
// 返回一个表示不与几何B相交的几何A的那部分的几何。可以将其视为GeometryA-ST_Intersection（A，B）。如果A完全包含在B中，则返回一个空的几何集合。
func (r *PostGis) Difference(geometry interface{}, geometry2 interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_Difference('%s','%s')", geometry, geometry2)
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// Union 合并
// http://postgis.net/docs/manual-3.0/ST_Union.html
func (r *PostGis) Union(geometry interface{}, geometry2 interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_Union(ST_AsText('%s'),ST_AsText('%s'))", geometry, geometry2) // 需要AsText否则会有问题
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)

	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	return
}

// MemUnion 合并
func (r *PostGis) MemUnion(geometry interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_MemUnion('%s')", geometry) // 需要AsText否则会有问题
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)

	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// Distance 距离
func (r *PostGis) Distance(geometry interface{}, geometry2 interface{}, returnType string) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_Distance('%s','%s')", geometry, geometry2) // 需要AsText否则会有问题
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	return
}

// DWithin DWithin
// 如果几何之间的距离在指定范围之内，则返回true。
// 对于几何：距离以几何的空间参照系定义的单位指定。为了使此功能有意义，源几何必须都具有相同的坐标投影，并具有相同的SRID。
// 对于地理单位，单位为米，默认将测量值设置为use_spheroid= true，以便进行更快速的检查；如果设置use_spheroid为false，则沿球面进行测量。
func (r *PostGis) DWithin(geometry interface{}, geometry2 interface{}, returnType string, distance float64) (result interface{}, err error) {
	//
	tmpStr := fmt.Sprintf("ST_DWithin('%s','%s','%f')", geometry, geometry2, distance) // 需要AsText否则会有问题
	SQL := r.getReturnType(returnType, tmpStr)
	SQL = fmt.Sprintf("SELECT %s;", SQL)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)

	return
}

// Equals Equals
func (r *PostGis) Equals(geometry interface{}, geometry2 interface{}) (equals bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Equals(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&equals)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Disjoint Disjoint
func (r *PostGis) Disjoint(geometry interface{}, geometry2 interface{}) (disjoint bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Disjoint(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&disjoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Intersects 相交
func (r *PostGis) Intersects(geometry interface{}, geometry2 interface{}) (intersects bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Intersects(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&intersects)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Touches Touches
func (r *PostGis) Touches(geometry interface{}, geometry2 interface{}) (touches bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Touches(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&touches)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Crosses Crosses
func (r *PostGis) Crosses(geometry interface{}, geometry2 interface{}) (crosses bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Crosses((%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&crosses)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Within Within
func (r *PostGis) Within(geometry interface{}, geometry2 interface{}) (within bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Within(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&within)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Overlaps Overlaps
func (r *PostGis) Overlaps(geometry interface{}, geometry2 interface{}) (overlaps bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Overlaps(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&overlaps)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Contains Contains
func (r *PostGis) Contains(geometry interface{}, geometry2 interface{}) (contains bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Contains(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&contains)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Covers Covers
func (r *PostGis) Covers(geometry interface{}, geometry2 interface{}) (covers bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Covers(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&covers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// CoveredBy CoveredBy
func (r *PostGis) CoveredBy(geometry interface{}, geometry2 interface{}) (coveredBy bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_CoveredBy(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&coveredBy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Relate Relate
func (r *PostGis) Relate(geometry interface{}, geometry2 interface{}) (relate bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Relate(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&relate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// AsText AsText
func (r *PostGis) AsText(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_AsText(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// AsBinary AsBinary
func (r *PostGis) AsBinary(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_AsBinary(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// SRID SRID
func (r *PostGis) SRID(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_SRID(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Dimension Dimension
func (r *PostGis) Dimension(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Dimension(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// Envelope Envelope
func (r *PostGis) Envelope(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Envelope(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// IsEmpty IsEmpty
func (r *PostGis) IsEmpty(geometry interface{}) (empty bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_AsText(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&empty)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// IsSimple IsSimple
func (r *PostGis) IsSimple(geometry interface{}) (simple bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_IsSimple(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&simple)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// IsClosed IsClosed
func (r *PostGis) IsClosed(geometry interface{}) (closed bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_IsClosed(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&closed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// IsRing IsRing
func (r *PostGis) IsRing(geometry interface{}) (ring bool, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_IsRing(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&ring)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false, err
	}
	return
}

// NumGeometries NumGeometries
func (r *PostGis) NumGeometries(geometry interface{}) (num int64, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_NumGeometries(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// GeometryN GeometryN
func (r *PostGis) GeometryN(geometry interface{}, index int64) (num int64, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_GeometryN(%s, %d);", geometry, index)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// NumPoints NumPoints
func (r *PostGis) NumPoints(geometry interface{}) (num int64, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_NumPoints(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// PointN PointN
func (r *PostGis) PointN(geometry interface{}, index int64) (num int64, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_PointN(%s, %d);", geometry, index)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// ExteriorRing ExteriorRing
func (r *PostGis) ExteriorRing(geometry interface{}) (num int64, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_ExteriorRing(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// NumInteriorRings NumInteriorRings
func (r *PostGis) NumInteriorRings(geometry interface{}) (num int64, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_NumInteriorRings(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// NumInteriorRing NumInteriorRing
func (r *PostGis) NumInteriorRing(geometry interface{}) (num int64, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_NumInteriorRing(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// InteriorRingN InteriorRingN
func (r *PostGis) InteriorRingN(geometry interface{}, index int64) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_InteriorRingN(%s, %d);", geometry, index)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// EndPoint EndPoint
func (r *PostGis) EndPoint(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_EndPoint(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// StartPoint StartPoint
func (r *PostGis) StartPoint(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_EndPoint(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// X X
func (r *PostGis) X(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_X(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// Y Y
func (r *PostGis) Y(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_EndPoint(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// M M
func (r *PostGis) M(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_EndPoint(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}

// Z Z
func (r *PostGis) Z(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_EndPoint(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}
