import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { BankAccount } from './models/bank-account.model';
import { BankAccountController } from './controllers/bank-account/bank-account.controller';
import { ConsoleModule } from 'nestjs-console';
import { FixtureCommand } from './fixtures/fixtures.command';
import { PixController } from './controllers/pix/pix.controller';
import { Pix } from './models/pix.model';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ConfigModule.forRoot(),
    ConsoleModule,
    TypeOrmModule.forRoot({
      type: process.env.TYPEORM_CONNECTION as any,
      host: process.env.TYPEORM_HOST,
      port: parseInt(process.env.TYPEORM_POSRT),
      username: process.env.TYPEORM_USERNAME,
      password: process.env.TYPEORM_PASSWORD,
      database: process.env.TYPEORM_DATABASE,
      entities: [BankAccount, Pix]               //available entities
    }),
    TypeOrmModule.forFeature([BankAccount, Pix]), //what entities I'll use
    ClientsModule.register([
      {
        name: 'CODEPIX_PACKAGE',
        transport: Transport.GRPC,
        options: {
          url: process.env.GRPC_URL,
          package: 'github.com.diegoclair.codepix', // is the protofile package
          protoPath: [join(__dirname, 'protofiles/pix.proto')]
        }
      }
    ])
  ],
  controllers: [AppController, BankAccountController, PixController],
  providers: [AppService, FixtureCommand],
})
export class AppModule {}
