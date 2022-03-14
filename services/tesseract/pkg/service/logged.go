package service

import (
	"context"

	log "github.com/sirupsen/logrus"

	pb "tesseract/pkg/service/pb"
)

type LoggedTesseractServer struct {
	logger *log.Logger
	next   pb.TesseractServer
}

func NewLoggedTesseractServer(logger *log.Logger, next pb.TesseractServer) *LoggedTesseractServer {
	return &LoggedTesseractServer{logger: logger, next: next}
}

func (l *LoggedTesseractServer) Apply(ctx context.Context, request *pb.ApplyRequest) (response *pb.ApplyResponse, err error) {
	response, err = l.next.Apply(ctx, request)
	if err == nil {
		l.logger.WithFields(log.Fields{
			"Namespace": request.ID,
			"Name":      request.Name,
			"DNS":       request.DNS,
			"Image":     request.Image,
			"Port":      request.Port,
			"Scale":     request.Scale,
			"CPU":       request.CPU,
			"MemoryMB":  request.RAM,
		}).Info("Apply")
	} else {
		l.logger.WithFields(log.Fields{
			"Namespace": request.ID,
			"Name":      request.Name,
			"DNS":       request.DNS,
			"Image":     request.Image,
			"Port":      request.Port,
			"Scale":     request.Scale,
			"CPU":       request.CPU,
			"MemoryMB":  request.RAM,
			"error": err,
		}).Info("Apply")
	}
	return
}

func (l *LoggedTesseractServer) GetStatus(ctx context.Context, request *pb.GetStatusRequest) (response *pb.GetStatusResponse, err error) {
	response, err = l.next.GetStatus(ctx, request)
	if err == nil {
		l.logger.WithFields(log.Fields{
			"ID":     request.ID,
			"status": response.Status.String(),
			"status_error": response.Error,
		}).Info("GetStatus")
	} else {
		l.logger.WithFields(log.Fields{
			"ID":     request.ID,
			"error": err,
		}).Info("GetStatus")
	}
	return
}

func (l *LoggedTesseractServer) Delete(ctx context.Context, request *pb.DeleteRequest) (response *pb.DeleteResponse, err error) {
	response, err = l.next.Delete(ctx, request)
	l.logger.WithFields(log.Fields{
		"ID": request.ID,
		"error": err,
	}).Info("Delete")
	return
}


