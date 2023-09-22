export namespace config {
	
	export class Config {
	    // Go type: struct { Version string "json:\"version\""; Build string "json:\"build\"" }
	    AppInfo: any;
	    // Go type: struct { Protections bool "json:\"protections\""; ClientPath string "json:\"clientPath\""; Server struct { AkiServerPath string "json:\"akiServerPath\""; MtgaServerPath string "json:\"mtgaServerPath\""; MtgaServerAddress string "json:\"mtgaServerAddress\""; AkiServerAddress string "json:\"akiServerAddress\"" } "json:\"server\""; Language string "json:\"language\""; Theme string "json:\"theme\""; LastProfile string "json:\"lastProfile\"" }
	    UserSettings: any;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.AppInfo = this.convertValues(source["AppInfo"], Object);
	        this.UserSettings = this.convertValues(source["UserSettings"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace profile {
	
	export class ProfileInfo {
	    id: string;
	    username: string;
	    edition: string;
	    side: string;
	    experience: number;
	    level: number;
	    head: string;
	    SavageLockTime: number;
	    usdAmount: number;
	    rubAmount: number;
	    eurAmount: number;
	
	    static createFrom(source: any = {}) {
	        return new ProfileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.edition = source["edition"];
	        this.side = source["side"];
	        this.experience = source["experience"];
	        this.level = source["level"];
	        this.head = source["head"];
	        this.SavageLockTime = source["SavageLockTime"];
	        this.usdAmount = source["usdAmount"];
	        this.rubAmount = source["rubAmount"];
	        this.eurAmount = source["eurAmount"];
	    }
	}
	export class ProfileList {
	    profiles: ProfileInfo[];
	
	    static createFrom(source: any = {}) {
	        return new ProfileList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.profiles = this.convertValues(source["profiles"], ProfileInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

