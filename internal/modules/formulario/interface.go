package formulario

// AvaluoCreator define el contrato necesario para crear un avalúo automático
// sin depender del paquete 'avaluo', evitando ciclos.
type AvaluoCreator interface {
	CrearAvaluoAutomatico(idFormulario uint) (any, error)
}
