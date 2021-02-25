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

// Area Area
func (r *PostGis) Area(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Area();", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
	return
}

// Length Length
func (r *PostGis) Length(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_Length();", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// PointOnSurface PointOnSurface
func (r *PostGis) PointOnSurface(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_PointOnSurface(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// Boundary Boundary
func (r *PostGis) Boundary(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_Boundary(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// Buffer Buffer
func (r *PostGis) Buffer(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_Buffer(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// ConvexHull ConvexHull
func (r *PostGis) ConvexHull(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_ConvexHull(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// Intersection Intersection
func (r *PostGis) Intersection(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_Intersection(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// ShiftLongitude ShiftLongitude
func (r *PostGis) ShiftLongitude(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_Shift_Longitude(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// SymDifference SymDifference
func (r *PostGis) SymDifference(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_SymDifference(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// Difference Difference
func (r *PostGis) Difference(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_Difference(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// Union Union
func (r *PostGis) Union(geometry interface{}) (result interface{}, err error) {
	//

	SQL := fmt.Sprintf("SELECT ST_Union(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// MemUnion MemUnion
func (r *PostGis) MemUnion(geometry interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_MemUnion(%s);", geometry)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// Distance Distance
func (r *PostGis) Distance(geometry interface{}, geometry2 interface{}) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_Distance(%s, %s);", geometry, geometry2)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

	return
}

// DWithin DWithin
func (r *PostGis) DWithin(geometry interface{}, geometry2 interface{}, distance float64) (result interface{}, err error) {
	//
	SQL := fmt.Sprintf("SELECT ST_DWithin(%s, %s, %b);", geometry, geometry2, distance)
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)

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

// Intersects Intersects
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
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
	err = r.Pool.QueryRow(context.Background(), SQL).Scan(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return -1, err
	}
	return
}
