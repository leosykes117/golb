describe('Pruebas para la autenticación y registro del sistema', () => {
  
  it('Iniciando sesión con un usuario existente', () => {
    cy.visit('/signin')

    cy.get(':nth-child(1) > .el-form-item__content > .el-input > .el-input__inner')
      .type('leo@gmail.com')
    
    cy.get('.is-required > .el-form-item__content > .el-input > .el-input__inner')
      .type('passw321')

    cy.get('.el-input__suffix-inner > .el-input__icon').eq(1).click()

    cy.get('.el-button').click()
  })

  it('Iniciando sesión con un usuario no existente', () => {
    cy.visit('/signin')

    cy.get(':nth-child(1) > .el-form-item__content > .el-input > .el-input__inner')
      .type('luis@gmail.com')
    
    cy.get('.is-required > .el-form-item__content > .el-input > .el-input__inner')
      .type('1234567890')

    cy.get('.el-button').click()

    cy.get('p.el-message__content').contains('Correo electrónico o contraseña incorrecta')
  })

  it('Registrando un nuevo usuario', () => {
    cy.visit('/signup')

    cy.get('h1').contains("Registrate")

    cy.get('input.el-input__inner').eq(0).type('Luis')
    cy.get('input.el-input__inner').eq(1).type('Ortiz')
    cy.get('input.el-input__inner').eq(2).type('luis@gmail.com')
    cy.get('input.el-input__inner').eq(3).type('1234567890')
    cy.get('input.el-input__inner').eq(4).type('1234567890')
    cy.get(':nth-child(1) > .el-switch__core').click()
  })
})