import {MigrationInterface, QueryRunner, Table, TableForeignKey} from "typeorm";

export class CreatePixTable1612745703956 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.createTable(
            new Table({
                name: 'pix',
                columns: [
                    {
                        name: 'id',
                        type: 'uuid',
                        isPrimary: true
                    },
                    {
                        name: 'keyType',
                        type: 'varchar'
                    },
                    {
                        name: 'key',
                        type: 'varchar'
                    },
                    {
                        name: 'bank_account_id',
                        type: 'uuid'
                    },
                    {
                        name: 'created_at',
                        type: 'timestamp',
                        default: 'CURRENT_TIMESTAMP'
                    },
                ],
            }),
        );

        await queryRunner.createForeignKey(
            'pix',
            new TableForeignKey({
                name: 'pix_bank_accounts_id_foreign_key',
                columnNames: ['bank_account_id'],
                referencedColumnNames: ['id'],
                referencedTableName: 'bank_accounts',
            }),
        );
    };

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.dropForeignKey(
            'pix',
            'pix_bank_accounts_id_foreign_key',
        );
        await queryRunner.dropTable('pix');
    };

}
