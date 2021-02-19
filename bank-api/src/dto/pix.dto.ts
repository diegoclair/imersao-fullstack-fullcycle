import { IsIn, IsNotEmpty, IsString } from "class-validator"
import { PixKeyType } from "src/models/pix.model"

export class PixDto {
    @IsString()
    @IsNotEmpty()
    readonly key: string

    @IsString()
    @IsIn(Object.values(PixKeyType))
    @IsNotEmpty()
    readonly key_type: PixKeyType
}