import { NestFactory } from '@nestjs/core';
import { Transport } from '@nestjs/microservices';
import { AppModule } from './app.module';
import { ModelNotFoundExceptionFilter } from './exception-filters/model-not-found.exception-filter';

async function bootstrap() {
  //{cors: true} is to open for all hosts
  const app = await NestFactory.create(AppModule, {cors: true});
  app.setGlobalPrefix('api')
  app.useGlobalFilters(new ModelNotFoundExceptionFilter());
  app.connectMicroservice({
    transport: Transport.KAFKA,
    options: {
      client: {
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
  });  
  app.startAllMicroservices();
  await app.listen(3000);
}
bootstrap();
