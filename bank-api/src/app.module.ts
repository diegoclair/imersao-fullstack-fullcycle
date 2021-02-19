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
import { TransactionController } from './controllers/transaction/transaction.controller';
import { Transaction } from './models/transaction.model';
import { TransactionSubscriber } from './subscribers/transaction-subscriber/transaaction-subscriber.service';

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
      entities: [BankAccount, Pix, Transaction]               //available entities
    }),
    TypeOrmModule.forFeature([BankAccount, Pix, Transaction]), //what entities I'll use
    ClientsModule.register([
      {
        name: 'GRPC_CODEPIX_PACKAGE',
        transport: Transport.GRPC,
        options: {
          url: process.env.GRPC_URL,
          package: 'github.com.diegoclair.codepix', // is the protofile package
          protoPath: [join(__dirname, 'protofiles/pix.proto')]
        }
      }
    ]),
    ClientsModule.register([
      {
        name: 'KAFKA_CODEPIX_TRANSACTION_SERVICE',
        transport: Transport.KAFKA,
        options: {
          client: {
            clientId: process.env.KAFKA_CLIENT_ID,
            brokers: [process.env.KAFKA_BROKER]
          },
          consumer: {
            //for dev we generate a random groupID, in production, needs to be a fixed one
            groupId: !process.env.KAFKA_CONSUMER_GROUP_ID ||
              process.env.KAFKA_CONSUMER_GROUP_ID === ''
                ? 'my-consumer-' + Math.random()
                : process.env.KAFKA_CONSUMER_GROUP_ID,
          }
        },
      },
    ]),
  ],
  controllers: [AppController, BankAccountController, PixController, TransactionController],
  providers: [AppService, FixtureCommand, TransactionSubscriber],
})
export class AppModule {}
