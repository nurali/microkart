package middleware

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

const RequestIDHeader = "X-Request-Id"

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func Chain(mw ...Middleware) Middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		// weav chain
		// last_mw gets final_handler as next, second_last_mw gets last_mw as next and it goes on
		// the chain would be first_mw get second_mw as next, second get third as next and finally last get final as next
		last := final
		for i := len(mw) - 1; i >= 0; i-- {
			last = mw[i](last)
		}

		// return func which call very first_mw in chain, which calls next and finally the last_mw calls final handler
		return func(rw http.ResponseWriter, r *http.Request) {
			last(rw, r)
		}
	}
}

func RequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id, _ := uuid.NewV4()
		r.Header.Set(RequestIDHeader, id.String())
		next(rw, r)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Debugf("got request, ReqId:%s", rw.Header().Get(RequestIDHeader))
		defer log.Debugf("done request, ReqId:%s", rw.Header().Get(RequestIDHeader))
		next(rw, r)
	}
}

func Recover(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("recovered from panic, err:%v", r)
				rw.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next(rw, r)
	}
}
