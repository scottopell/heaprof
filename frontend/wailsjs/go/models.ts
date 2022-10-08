export namespace main {
	
	export class PprofLoadedData {
	    fileName: string;
	
	    static createFrom(source: any = {}) {
	        return new PprofLoadedData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	    }
	}

}

