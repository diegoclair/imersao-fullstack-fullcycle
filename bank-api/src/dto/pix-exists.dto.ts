import { IsIn, IsNotEmpty, IsString } from "class-validator"
import { PixKeyType } from "src/models/pix.model"

export class PixExistsDto {
    @IsString()
    @IsNotEmpty()
    readonly key: string

    @IsString()
    @IsIn(Object.values(PixKeyType))
    @IsNotEmpty()
    readonly keyType: PixKeyType
}