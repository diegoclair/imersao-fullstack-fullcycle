import { command } from "commander";
import { Command, Console } from "nestjs-console";
import { getConnection, getRepository } from "typeorm";
import fixtures from "./fixtures";
import * as chalk from 'chalk'

@Console()
export class FixtureCommand{
    
    @Command({
        command: 'fixtures',
        description: 'Seed data in database'
    })
    async command(){
        await this.runMigrations();
        for(const fixture of fixtures){
            await this.createInDatabase(fixture.model, fixture.fields)
        }

        console.log(chalk.green('Data generated'));
        
    }

    async runMigrations() {
        const conn = getConnection('default');
        for(const migration of conn.migrations.reverse()){
            await conn.undoLastMigration();
        }
    }

    async createInDatabase(model: any, data: any) {
        const repository = this.getRepository(model);
        const obj = repository.create(data);
        await repository.save(obj)
    }

    getRepository(model: any){
        const conn = getConnection('default');
        return conn.getRepository(model)
    }
}
