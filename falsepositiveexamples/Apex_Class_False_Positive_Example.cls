//asist-ignore-begin:[ApexClassWithoutSharing,ApexClassNoSharing] API is designed to allow access to all documents therefore needs to be without sharing.
public without sharing class xyz{
// asist-ignore-end
    
    public with sharing class abc{
    //asist-ignore-begin:[InsecureCryptoAlgorithm]
    Crypto.generateDigest('SHA1', Blob.valueOf(xyz))
    // asist-ignore-end 
    }
}



