import {HttpEvent,HttpHandler,HttpInterceptor,HttpRequest, HttpResponse, HttpErrorResponse} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import { SpinnerServiceService } from './services/spinner-service.service';
import { tap, map } from 'rxjs/operators';

@Injectable()
export class CustomHttpInterceptor implements HttpInterceptor{

    constructor(private spinnerService:SpinnerServiceService){

    }

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        console.log("HOLA DESDE AQUI")
        this.spinnerService.requestStarted();
        return this.handler(next,req);
        //return next.handle(req);
    }

    handler(next,req){
        return next.handle(req)
            .pipe(
                tap(
                    (event)=>{
                        if(event instanceof HttpResponse){/*Esta parte recive nuestra respuesta */
                            this.spinnerService.requestEnded();
                        }
                    },
                    (error:HttpErrorResponse)=>{
                        this.spinnerService.resetSpinner();
                        throw error;
                    }
                ),
            );
    }

}