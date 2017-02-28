package cpuinfo

import (
	"github.com/emicklei/go-restful"
    cgl_cpuinfo "cgolib/cpuinfo"
)


// FIXME(Shaohe Feng), Maybe we need a midleware layer here, to translate
// the pqos data to API data. Maybe we do not need to expose all fields of
// cpuinfo from pqos.
type Cpuinfo struct {

}

type CpuinfoResource struct {

}

func (cpuinfo CpuinfoResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
    // FIXME(Shaohe Feng)  here is hard code for v1, need refactor.
	ws.
		Path("/v1").
		Doc("Show the cupinfo of a host.").
        // FIXME(Shaohe Feng) just need to support json.
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/cpuinfo").To(cpuinfo.getCpuinfo).
		// docs
		Doc("get cpuinfo").
		Operation("getCpuinfo").
		Writes(Cpuinfo{})) // on the response

	container.Add(ws)
}


// FIXME(Shaohe Feng) localhost:8081
// GET http://localhost:8081/v1/cpuinfo
func (cpuinfo CpuinfoResource) getCpuinfo(request *restful.Request,
                                          response *restful.Response) {
    pq, _ := cgl_cpuinfo.GetCpuInfo()
	response.WriteEntity(pq)
}
