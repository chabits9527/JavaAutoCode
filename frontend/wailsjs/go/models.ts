export namespace entity {
	
	export class DatabaseConfig {
	    id: number;
	    name: string;
	    dbType: string;
	    host: string;
	    port: string;
	    username: string;
	    password: string;
	    database: string;
	    params: string;
	
	    static createFrom(source: any = {}) {
	        return new DatabaseConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.dbType = source["dbType"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.database = source["database"];
	        this.params = source["params"];
	    }
	}
	export class PageInfo {
	    total: number;
	    pageNum: number;
	    pageSize: number;
	    list: any;
	
	    static createFrom(source: any = {}) {
	        return new PageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.pageNum = source["pageNum"];
	        this.pageSize = source["pageSize"];
	        this.list = source["list"];
	    }
	}

}

