import { ArgumentsHost, Catch, ExceptionFilter } from "@nestjs/common";
import { Response } from "express";
import { EntityNotFoundError } from 'typeorm/error/EntityNotFoundError';

@Catch(EntityNotFoundError)
export class ModelNotFoundExceptionFilter implements ExceptionFilter {
    
    catch(exception: EntityNotFoundError, host: ArgumentsHost){
        const ctx = host.switchToHttp();
        const response = ctx.getResponse<Response>();
        return response.status(404).json({
            error: {
                error: 'Not Found',
                message: exception.message,

            },
        });
    }
}