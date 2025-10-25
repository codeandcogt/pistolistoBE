package avaluo

// AvaluoCreator define solo el contrato que Formulario necesita conocer
type AvaluoCreator interface {
	CrearAvaluoAutomatico(idFormulario uint, idUsuario int) error
}
