package pkg.proto2;

import static pkg.proto2.Testdata2.ExtensionMessage;
import static pkg.proto2.Testdata2.OuterMessage;

public class Proto2 {
    public static void main() {
        //- @OuterMessage ref JavaOuterMessage
        OuterMessage msg =
            OuterMessage.newBuilder().build();
        //- @ExtensionMessage ref JavaExtensionMessage
        //- @extension ref JavaExtensionField
        msg.hasExtension(ExtensionMessage.extension);
    }

    //- OuterMessage generates JavaOuterMessage

    //- ExtensionMessage generates JavaExtensionMessage
    //- ExtensionField generates JavaExtensionField
}
