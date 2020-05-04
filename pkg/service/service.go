package service


import (
	"context"
	dmap "github.com/vitush/go-grpc-dg-poc/pkg/api"
)

//
//import (
//	"context"
//	"database/sql"
//	"fmt"
//	"time"
//
//	"github.com/golang/protobuf/ptypes"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//
//	dmap "github.com/vitush/go-grpc-dg-poc/pkg/api"
//)
//const (
//	// apiVersion is version of API is provided by server
//	apiVersion = "v1"
//)
//
//// toDoServiceServer is implementation of dmap.ToDoServiceServer proto interface
//type toDoServiceServer struct {
//	db *sql.DB
//}
//
//// NewToDoServiceServer creates ToDo service
//func NewToDoServiceServer(db *sql.DB) dmap.ToDoServiceServer {
//	return &toDoServiceServer{db: db}
//}
//
//// checkAPI checks if the API version requested by client is supported by server
//func (s *toDoServiceServer) checkAPI(api string) error {
//	// API version is "" means use current version of the service
//	if len(api) > 0 {
//		if apiVersion != api {
//			return status.Errorf(codes.Unimplemented,
//				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
//		}
//	}
//	return nil
//}
//
//// connect returns SQL database connection from the pool
//func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
//	c, err := s.db.Conn(ctx)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
//	}
//	return c, nil
//}
//
//// Create new todo task
//func (s *toDoServiceServer) Create(ctx context.Context, req *dmap.CreateRequest) (*dmap.CreateResponse, error) {
//	// check if the API version requested by client is supported by server
//	if err := s.checkAPI(req.Api); err != nil {
//		return nil, err
//	}
//
//	// get SQL connection from pool
//	c, err := s.connect(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
//	}
//
//	// insert ToDo entity data
//	res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
//		req.ToDo.Title, req.ToDo.Description, reminder)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
//	}
//
//	// get ID of creates ToDo
//	id, err := res.LastInsertId()
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
//	}
//
//	return &dmap.CreateResponse{
//		Api: apiVersion,
//		Id:  id,
//	}, nil
//}
//
//// Read todo task
//func (s *toDoServiceServer) Read(ctx context.Context, req *dmap.ReadRequest) (*dmap.ReadResponse, error) {
//	// check if the API version requested by client is supported by server
//	if err := s.checkAPI(req.Api); err != nil {
//		return nil, err
//	}
//
//	// get SQL connection from pool
//	c, err := s.connect(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	// query ToDo by ID
//	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",
//		req.Id)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
//	}
//	defer rows.Close()
//
//	if !rows.Next() {
//		if err := rows.Err(); err != nil {
//			return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
//		}
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
//			req.Id))
//	}
//
//	// get ToDo data
//	var td dmap.ToDo
//	var reminder time.Time
//	if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
//		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
//	}
//	td.Reminder, err = ptypes.TimestampProto(reminder)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
//	}
//
//	if rows.Next() {
//		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%d'",
//			req.Id))
//	}
//
//	return &dmap.ReadResponse{
//		Api:  apiVersion,
//		ToDo: &td,
//	}, nil
//
//}
//
//// Update todo task
//func (s *toDoServiceServer) Update(ctx context.Context, req *dmap.UpdateRequest) (*dmap.UpdateResponse, error) {
//	// check if the API version requested by client is supported by server
//	if err := s.checkAPI(req.Api); err != nil {
//		return nil, err
//	}
//
//	// get SQL connection from pool
//	c, err := s.connect(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
//	}
//
//	// update ToDo
//	res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?",
//		req.ToDo.Title, req.ToDo.Description, reminder, req.ToDo.Id)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to update ToDo-> "+err.Error())
//	}
//
//	rows, err := res.RowsAffected()
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
//	}
//
//	if rows == 0 {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
//			req.ToDo.Id))
//	}
//
//	return &dmap.UpdateResponse{
//		Api:     apiVersion,
//		Updated: rows,
//	}, nil
//}
//
//// Delete todo task
//func (s *toDoServiceServer) Delete(ctx context.Context, req *dmap.DeleteRequest) (*dmap.DeleteResponse, error) {
//	// check if the API version requested by client is supported by server
//	if err := s.checkAPI(req.Api); err != nil {
//		return nil, err
//	}
//
//	// get SQL connection from pool
//	c, err := s.connect(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	// delete ToDo
//	res, err := c.ExecContext(ctx, "DELETE FROM ToDo WHERE `ID`=?", req.Id)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to delete ToDo-> "+err.Error())
//	}
//
//	rows, err := res.RowsAffected()
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
//	}
//
//	if rows == 0 {
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
//			req.Id))
//	}
//
//	return &dmap.DeleteResponse{
//		Api:     apiVersion,
//		Deleted: rows,
//	}, nil
//}
//
//// Read all todo tasks
//func (s *toDoServiceServer) ReadAll(ctx context.Context, req *dmap.ReadAllRequest) (*dmap.ReadAllResponse, error) {
//	// check if the API version requested by client is supported by server
//	if err := s.checkAPI(req.Api); err != nil {
//		return nil, err
//	}
//
//	// get SQL connection from pool
//	c, err := s.connect(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	// get ToDo list
//	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo")
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
//	}
//	defer rows.Close()
//
//	var reminder time.Time
//	list := []*dmap.ToDo{}
//	for rows.Next() {
//		td := new(dmap.ToDo)
//		if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
//			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
//		}
//		td.Reminder, err = ptypes.TimestampProto(reminder)
//		if err != nil {
//			return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
//		}
//		list = append(list, td)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
//	}
//
//	return &dmap.ReadAllResponse{
//		Api:   apiVersion,
//		ToDos: list,
//	}, nil
//}


// dMapServiceServer is implementation of dmap.DMapServiceServer proto interface
type dMapServiceServer struct {
	x *string
}

// NewToDoServiceServer creates dMapService service
func NewDMapServiceServer() dmap.DMapServiceServer{
	return &dMapServiceServer{x: nil}
}

// GetOrCreateMap
func (s *dMapServiceServer) GetOrCreateMap(ctx context.Context, req *dmap.GetOrCreateMapRequest) (*dmap.GetOrCreateMapResponse, error) {
	return &dmap.GetOrCreateMapResponse{
		MapObject:            nil,
	}, nil
}

// GetOrCreateMapRedis
func (s *dMapServiceServer) GetOrCreateMapRedis(ctx context.Context, req *dmap.GetOrCreateMapRequest) (*dmap.GetOrCreateMapResponse, error) {
	return &dmap.GetOrCreateMapResponse{
		MapObject:            nil,
	}, nil
}

// DmapGet
func (s *dMapServiceServer) DmapGet(ctx context.Context, req *dmap.DMapGetRequest) (*dmap.DMapGetResponse, error) {
	return &dmap.DMapGetResponse{
		Value:                "",
	}, nil
}

// DmapGet
func (s *dMapServiceServer) DmapGetRedis(ctx context.Context, req *dmap.DMapGetRequest) (*dmap.DMapGetResponse, error) {
	return &dmap.DMapGetResponse{
		Value:                "",
	}, nil
}

// DmapSet
func (s *dMapServiceServer) DmapSet(ctx context.Context, req *dmap.DMapSetRequest) (*dmap.DMapSetResponse, error) {
	return &dmap.DMapSetResponse{
		Status:               false,
	}, nil
}

// DmapSetRedis
func (s *dMapServiceServer) DmapSetRedis(ctx context.Context, req *dmap.DMapSetRequest) (*dmap.DMapSetResponse, error) {
	return &dmap.DMapSetResponse{
		Status:               false,
	}, nil
}
