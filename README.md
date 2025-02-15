# eCommerce-maqueta Ing: Junior 
 Tarea 2 Crear la maqueta inicial de la página web de la plataforma e-Commerce, que demuestre el dominio de HTML y CSS avanzados, incluyendo HTML5, CSS3 y el uso de frameworks como Bootstrap o Tailwind CSS.

E-Commerce Junior

Estructura del Proyecto
Use Docker es una plataforma diseñada para ayudar a los desarrolladores a crear, compartir y ejecutar aplicaciones en contenedores.
e-commerce-Junior-front/: Directorio que contiene el código fuente de la aplicación Angular para front.
e-commerce-Junior-backend/: Directorio que contiene el código fuente de la aplicación Go para Backend.
e-commerce-junior-db/mysqldata/: Directorio para almacenar los datos de MySQL para base datos.
e-commerce-junior-mongodb/data/: Directorio para almacenar los datos de MongoDB para usar en aws solo inegre para futuro.

Creacion de login
![Captura de pantalla 2025-01-30 175231](https://github.com/user-attachments/assets/8db85bc3-6ea8-4eab-af01-7a1e6ba66383)
Validacion login
![Captura de pantalla 2025-01-30 180318](https://github.com/user-attachments/assets/193d3df6-50ce-4775-b40b-ee8cd29ca6ff)
Creacion de Registro 
![Captura de pantalla 2025-01-30 180213](https://github.com/user-attachments/assets/9b72dda5-f2ab-432e-b70c-a583a5bc75dc)
Validacion de Registro
![Captura de pantalla 2025-01-30 180355](https://github.com/user-attachments/assets/92398724-7f5e-498e-a2b1-95f7b7eeef3b)

Formularios con validación
![Captura de pantalla 2025-01-31 084011](https://github.com/user-attachments/assets/b45d67fc-acd9-4582-88ac-a312a2ca85bd)

Formulario Video
![Captura de pantalla 2025-01-30 234118](https://github.com/user-attachments/assets/b6faa740-c999-4046-8d33-3a661194bb79)
Formulario Audio Y canvas.
![Captura de pantalla 2025-01-31 085556](https://github.com/user-attachments/assets/1ba6cbde-8767-4d73-90f1-c420596820a0)

Codigo html5 y css 
![Captura de pantalla 2025-01-31 085717](https://github.com/user-attachments/assets/737af18c-6c55-4d9e-8f9c-61e854e82189)
![Captura de pantalla 2025-01-31 085754](https://github.com/user-attachments/assets/8571675b-9238-4751-ad05-f7db3944da17)
![Captura de pantalla 2025-01-31 085857](https://github.com/user-attachments/assets/131b2fb5-9da8-4263-a4f9-f228a703ba8c)


 
http://localhost:4200/login
 
Validador de Usuario 
<div class="login-container">
    <div class="card card-container">
        <div class="col-md-12">
            <h3 class="text-center">Ecommerce Biu Inicia Sesión para comenzar a comprar</h3>
            <img id="profile-img" src="//ssl.gstatic.com/accounts/ui/avatar_2x.png" class="profile-img-card" />
            <form name="form" (ngSubmit)="onSubmit(f)" #f="ngForm" novalidate>
                <div class="form-group">
                    <label for="username">Usuario</label>
                    <input type="text" class="form-control" name="username" [(ngModel)]="form.username" required
                        #username="ngModel" [ngClass]="{ 'is-invalid': f.submitted && username.errors }" />
                    <div *ngIf="username.errors && f.submitted" class="invalid-feedback">
                        El usuario es obligatorio
                    </div>
                </div>
                <div class="form-group">
                    <label for="password">Contraseña</label>
                    <input type="password" class="form-control" name="password" [(ngModel)]="form.password" required
                        minlength="6" #password="ngModel"
                        [ngClass]="{ 'is-invalid': f.submitted && password.errors }" />
                    <div *ngIf="password.errors && f.submitted" class="invalid-feedback">
                        <div *ngIf="password.errors['required']">La contraseña es obligatoria</div>
                        <div *ngIf="password.errors['minlength']">
                            La contraseña debe tener al menos 6 caracteres
                        </div>
                    </div>
                </div>
                <div class="form-group">
                    <button class="btn btn-primary btn-block">
                        Ingresar
                    </button>
                </div>
                <div class="form-group social-netwoks">
                    <button type="button" class="btn btn-sm btn-google" (click)="loginWithGoogle()">
                        <i class="bi bi-google"></i>
                    </button>
                    <button type="button" class="btn btn-sm btn-facebook" (click)="loginWithFacebook()">
                        <i class="bi bi-facebook"></i>
                    </button>
                </div>
                <div class="form-group">>
                   <a [routerLink]="[ '/register' ]" routerLinkActive="active">Registrarse</a>
                </div>
                <div class="form-group">
                    <div *ngIf="f.submitted && isLoginFailed" class="alert alert-danger" role="alert">
                        Error: {{ errorMessage }}
                    </div>
                </div>
            </form>
        </div>
    </div>
</div>
 

 
http://localhost:4200/register
 

<div class="login-container">
    <div class="card card-container">
        <div class="col-md-12">
            <h3 class="text-center">Ingresa los datos solicitados para registrarse</h3>
            <img id="profile-img" src="//ssl.gstatic.com/accounts/ui/avatar_2x.png" class="profile-img-card" />

            <form [formGroup]="userForm" (ngSubmit)="onSubmit()" class="needs-validation" novalidate>

                <div class="row">

                    <div class="col-md-12 mb-3">
                        <label for="username" class="form-label">Nombre de Usuario:</label>
                        <input id="username" formControlName="username" class="form-control"
                            [ngClass]="{'is-invalid': userForm.get('username')?.invalid && userForm.get('username')?.touched}"
                            formControlName="username">
                        <div *ngIf="userForm.get('username')?.invalid && userForm.get('username')?.touched"
                            class="invalid-feedback">
                            El nombre de usuario es requerido.
                        </div>
                    </div>

                    <div class="col-md-12 mb-3">
                        <label for="email" class="form-label">Correo Electrónico:</label>
                        <input id="email" formControlName="email" class="form-control"
                            [ngClass]="{'is-invalid': userForm.get('email')?.invalid && userForm.get('email')?.touched}"
                            formControlName="email">
                        <div *ngIf="userForm.get('email')?.invalid && userForm.get('email')?.touched"
                            class="invalid-feedback">
                            Debe ingresar un correo electrónico válido
                        </div>
                    </div>

                    <div class="col-md-12 mb-3">
                        <label for="avatar" class="form-label">Avatar:</label>
                        <input id="avatar" formControlName="avatar" class="form-control"
                            [ngClass]="{'is-invalid': userForm.get('avatar')?.invalid && userForm.get('avatar')?.touched}"
                            formControlName="avatar">
                        <div *ngIf="userForm.get('avatar')?.invalid && userForm.get('avatar')?.touched"
                            class="invalid-feedback">
                            El avatar de usuario es requerido.
                        </div>
                    </div>

                    <div class="col-md-12 mb-3">
                        <label for="password" class=" form-label">Contraseña:</label>
                        <input id="password" type="password" formControlName="password" class="form-control"
                            [ngClass]="{'is-invalid': userForm.get('password')?.invalid && userForm.get('password')?.touched}"
                            formControlName="password">
                        <div *ngIf=" userForm.get('password')?.invalid && userForm.get('password')?.touched"
                            class="invalid-feedback">
                            La contraseña es obligatoria.
                        </div>
                    </div>

                    <div class="col-md-12 mb-3">
                        <label for="confirmPassword" class=" form-label">Confirmar Contraseña:</label>
                        <input id="confirmPassword" type="password" formControlName="confirmPassword" class="form-control"
                            [ngClass]="{'is-invalid': userForm.get('confirmPassword')?.invalid && userForm.get('confirmPassword')?.touched}"
                            formControlName="confirmPassword">
                        <div *ngIf="userForm.get('confirmPassword')?.invalid &&
                            userForm.get('confirmPassword')?.touched" class="invalid-feedback">
                            Debe ingresar nuevamente la contraseña
                        </div>
                    </div>

                    <div *ngIf="userForm.errors?.['mismatch']" class="invalid-feedback" style="display: block;">
                        Las contraseñas no coinciden
                    </div>

                    <div class="form-group">>
                      <a [routerLink]="[ '/login' ]" routerLinkActive="active">Volver a Login</a>
                   </div>
                    <div class="d-flex justify-content-end">
                        <button type="submit" class="btn d-block w-25 btn-success"
                            [disabled]="userForm.invalid">Guardar</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</div>

Para Iniciar el login cree la Base Datos MYSQL 
 ![image](https://github.com/user-attachments/assets/47f32e03-f6b6-45df-8de2-8e32172973c1)

 
![image](https://github.com/user-attachments/assets/197c36fa-8873-4cd5-be43-c2fe8f5f939c)



Crear Listado de Producto Inventario
 
 ![image](https://github.com/user-attachments/assets/f609f867-0b6f-421d-b023-e2598e961d83)





