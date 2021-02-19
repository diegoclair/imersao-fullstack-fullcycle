import { IsNotEmpty, IsNumber, IsOptional, IsString, Min } from "class-validator";
import { PixKeyType } from "src/models/pix.model";

export class TransactionDto {
    @IsString()
    @IsNotEmpty()
    pix_key_to : string;
  
    @IsString()
    @IsNotEmpty()
    pix_key_type_to : PixKeyType;
  
    @IsString()
    @IsOptional()
    description : string = null;
  
    @IsNumber({ maxDecimalPlaces: 2 })
    @Min(0.01)
    @IsNotEmpty()
    readonly amount: number;
}