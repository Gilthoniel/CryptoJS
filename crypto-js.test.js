import './crypto-js'

window.crypto = {
  getRandomValues(array) {
    array.forEach((a, i) => {
      array[i] = Math.round(Math.random() * 100000)
    });
  }
};

const MOCK_KEYS = [{
  private: "9d61b19deffd5a60ba844af492ec2cc44449c5697b326919703bac031cae7f60",
  public: "d75a980182b10ab7d54bfed3c964073a0ee172f3daa62325af021a68f707511a",
  msg: "",
  signature: "e5564300c360ac729086e2cc806e828a84877f1eb8e5d974d873e065224901555fb8821590a33bacc61e39701cf9b46bd25bf5f0595bbe24655141438e7a100b"
}, {
  private: "4ccd089b28ff96da9db6c346ec114e0f5b8a319f35aba624da8cf6ed4fb8a6fb",
  public: "3d4017c3e843895a92b70aa74d1b7ebc9c982ccf2ec4968cc0cd55f12af4660c",
  msg: "72",
  signature: "92a009a9f0d4cab8720e820b5f642540a2b27b5416503f8fb3762223ebdb69da085ac1e43e15996e458f3613d0f11d8c387b2eaeb4302aeeb00d291612bb0c00"
}, {
  private: "c5aa8df43f9f837bedb7442f31dcb7b166d38535076f094b85ce3a2e0b4458f7",
  public: "fc51cd8e6218a1a38da47ed00230f0580816ed13ba3303ac5deb911548908025",
  msg: "af82",
  signature: "6291d657deec24024827e69c3abe01a30ce548a284743a445e3680d7db5ac3ac18ff9b538d16f290ae67f760984dc6594a7c15e9716ed28dc027beceea1ec40a"
}, {
  private: "f5e5767cf153319517630f226876b86c8160cc583bc013744c6bf255f5cc0ee5",
  public: "278117fc144c72340f67d0f2316e8386ceffbf2b2428c9c51fef7c597f1d426e",
  msg: "08b8b2b733424243760fe426a4b54908632110a66c2f6591eabd3345e3e4eb98fa6e264bf09efe12ee50f8f54e9f77b1e355f6c50544e23fb1433ddf73be84d879de7c0046dc4996d9e773f4bc9efe5738829adb26c81b37c93a1b270b20329d658675fc6ea534e0810a4432826bf58c941efb65d57a338bbd2e26640f89ffbc1a858efcb8550ee3a5e1998bd177e93a7363c344fe6b199ee5d02e82d522c4feba15452f80288a821a579116ec6dad2b3b310da903401aa62100ab5d1a36553e06203b33890cc9b832f79ef80560ccb9a39ce767967ed628c6ad573cb116dbefefd75499da96bd68a8a97b928a8bbc103b6621fcde2beca1231d206be6cd9ec7aff6f6c94fcd7204ed3455c68c83f4a41da4af2b74ef5c53f1d8ac70bdcb7ed185ce81bd84359d44254d95629e9855a94a7c1958d1f8ada5d0532ed8a5aa3fb2d17ba70eb6248e594e1a2297acbbb39d502f1a8c6eb6f1ce22b3de1a1f40cc24554119a831a9aad6079cad88425de6bde1a9187ebb6092cf67bf2b13fd65f27088d78b7e883c8759d2c4f5c65adb7553878ad575f9fad878e80a0c9ba63bcbcc2732e69485bbc9c90bfbd62481d9089beccf80cfe2df16a2cf65bd92dd597b0707e0917af48bbb75fed413d238f5555a7a569d80c3414a8d0859dc65a46128bab27af87a71314f318c782b23ebfe808b82b0ce26401d2e22f04d83d1255dc51addd3b75a2b1ae0784504df543af8969be3ea7082ff7fc9888c144da2af58429ec96031dbcad3dad9af0dcbaaaf268cb8fcffead94f3c7ca495e056a9b47acdb751fb73e666c6c655ade8297297d07ad1ba5e43f1bca32301651339e22904cc8c42f58c30c04aafdb038dda0847dd988dcda6f3bfd15c4b4c4525004aa06eeff8ca61783aacec57fb3d1f92b0fe2fd1a85f6724517b65e614ad6808d6f6ee34dff7310fdc82aebfd904b01e1dc54b2927094b2db68d6f903b68401adebf5a7e08d78ff4ef5d63653a65040cf9bfd4aca7984a74d37145986780fc0b16ac451649de6188a7dbdf191f64b5fc5e2ab47b57f7f7276cd419c17a3ca8e1b939ae49e488acba6b965610b5480109c8b17b80e1b7b750dfc7598d5d5011fd2dcc5600a32ef5b52a1ecc820e308aa342721aac0943bf6686b64b2579376504ccc493d97e6aed3fb0f9cd71a43dd497f01f17c0e2cb3797aa2a2f256656168e6c496afc5fb93246f6b1116398a346f1a641f3b041e989f7914f90cc2c7fff357876e506b50d334ba77c225bc307ba537152f3f1610e4eafe595f6d9d90d11faa933a15ef1369546868a7f3a45a96768d40fd9d03412c091c6315cf4fde7cb68606937380db2eaaa707b4c4185c32eddcdd306705e4dc1ffc872eeee475a64dfac86aba41c0618983f8741c5ef68d3a101e8a3b8cac60c905c15fc910840b94c00a0b9d0",
  signature: "0aab4c900501b3e24d7cdf4663326a3a87df5e4843b2cbdb67cbf6e460fec350aa5371b1508f9f4528ecea23c436d94b5e8fcd4f681e30a6ac00a9704a188a03"},
];

const MOCK_PUBLIC_KEYS = [
  "e5e23e58539a09d3211d8fa0fb3475d48655e0c06d83e93c8e6e7d16aa87c106",
  "036bf316e1ea6e7e99e0bb713419d16c0b6794bf9dc442cc4cf36c3f935e93cf",
  "9b54fdfbb39138f0a45a08e1721ab4eb257e0f246c93a272154d651454756290",
  "52c6a77c756ce5f3bef3414006c45556aeb7084c0f3e4467dd27107927912c51",
  "82e864d06230daa2e90a1ac24ec738bcd9e5458677e75f1e64a9b514a711d630",
  "588addc1fbd9c403868e0d0ac2003e7c7538d0a9154fbdb3cf772822f0ebc827"
];

const MOCK_AGG_RESULT = "688d8bfddea6f413a29e32c8ff0a1fb603beac0e719e7a395b3ad52ce1623246";

const MOCK_HASH = [{
  value: '',
  sha256: 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855',
  sha512: 'cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e'
},{
  value: 'This is a mock message used to test the hash functions',
  sha256: '6a88e4e7ffb58d3307d5ce3579cdc019e47b5990f03cd9b26356af81079944c1',
  sha512: 'b2d28852a63d7cb08b4f188f3306a86ef672b5c8b3ac9db6bcf9d496cb19129d8ccbc3de4a6084a3af182121ca2c12c7996bb6fcaec8494b067b6fdaa060074e'
}];

describe('crypto-js', () => {

  it('it should produce the same public key and signature', () => {
    expect.assertions(MOCK_KEYS.length * 3);

    MOCK_KEYS.forEach((mock) => {
      const pubKey = cryptoJS.generatePublicKey(mock.private);
      expect(pubKey).toBe(mock.public);

      const signature = cryptoJS.sign(mock.private, mock.msg);
      expect(signature).toBe(mock.signature);

      expect(cryptoJS.verify(pubKey, mock.msg, mock.private)).toBeTruthy();
    });
  });

  it('should aggregate correctly', () => {
    const agg = cryptoJS.aggregateKeys(MOCK_PUBLIC_KEYS);

    expect(agg).toBe(MOCK_AGG_RESULT);
  });

  it('should hash (sha256) correctly', () => {
    expect.assertions(MOCK_HASH.length);

    MOCK_HASH.forEach((mock) => {
      const hash = cryptoJS.sha256(mock.value);
      expect(hash).toBe(mock.sha256);
    });
  });

  it('should hash (sha512) correctly', () => {
    expect.assertions(MOCK_HASH.length);

    MOCK_HASH.forEach((mock) => {
      const hash = cryptoJS.sha512(mock.value);
      expect(hash).toBe(mock.sha512);
    });
  });

});